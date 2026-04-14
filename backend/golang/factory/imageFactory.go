package factory

import (
	"fmt"
	"log/slog"
	"my-go-api/cache"
	"my-go-api/storage"
)

type ImageFactory struct {
	storageService storage.StorageService
	redisService   cache.CacheService
}

func NewImageFactory(logger *slog.Logger) *ImageFactory {
	redisService := cache.NewRedisService(logger, "localhost:6379", "", 0)
	storageService := storage.NewFileStoreService(logger)
	return &ImageFactory{
		storageService: storageService,
		redisService:   redisService,
	}
}

func (imgFact *ImageFactory) GetRandomImageName() string {

	imgName, err := imgFact.redisService.GetImageName()
	if err == nil {
		return imgName
	}
	fmt.Println("Failed to get image name from redis: ", err)

	imgName, err = imgFact.storageService.GetRandomImageName()
	if err == nil {
		return imgName
	}
	fmt.Println("Failed to get image name from storage service: ", err)

	return ""
}

func (imgFact *ImageFactory) GetImage(imageName string) []byte {
	img, err := imgFact.storageService.GetImage(imageName)
	if err != nil {
		fmt.Println("Failed to get image content from storage service: ", err)
	}
	return img
}

func (imgFact *ImageFactory) RefreshCache() {
	imageNames := []string{}
	for i := 0; i < 30; i++ {
		imageName, err := imgFact.storageService.GetRandomImageName()
		if err != nil {
			fmt.Println(err)
		}
		imageNames = append(imageNames, imageName)
	}

	err := imgFact.redisService.CreateNewSet(imageNames)
	if err != nil {
		fmt.Println(err)
	}
}
