package service

import (
	"io"
	"io/fs"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var imagesCachedPerJob int = 100

type ImageService struct {
	logger    *slog.Logger
	Redis     RedisService
	Filenames *[]string
}

func NewImageService(logger *slog.Logger, r RedisService) *ImageService {
	return &ImageService{
		logger:    logger.With("imageService"),
		Redis:     r,
		Filenames: &[]string{},
	}
}

func (s ImageService) GetAllFilesFromDir() {
	output := []string{}
	fileTypes := []string{
		".jpg",
		".jpeg",
		".png",
		".gif",
		".bmp",
		".tiff",
		".webp",
	}
	imageRootDirectory := "/home/henry/Desktop/images/mountains"

	err := filepath.Walk(imageRootDirectory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			s.logger.Warn("failure accessing a path %q: %v\n", path, err)
			return err
		}

		for _, fileType := range fileTypes {
			if strings.HasSuffix(path, fileType) {
				output = append(output, path)
			}
		}

		if info.IsDir() {
			s.logger.Debug("visited dir: %q\n", path)
		}
		return nil
	})
	if err != nil {
		s.logger.Warn("error walking the path %q: %v\n", imageRootDirectory, err)
	}
	s.logger.Debug("output len %d\n", len(output))

	*s.Filenames = output
}

func (s ImageService) GetRandomImage() []byte {
	key := s.Redis.GetRandomKey()
	val, err := s.Redis.GetBytes(key)
	if err != nil {
		s.logger.Warn("Failed to retrieve image from cache, grabbing image file directly")
		_, val = s.GetRandomImageWithoutCache()
	}
	s.logger.Info("Returning image")
	return val
}

func (s ImageService) GetRandomImageWithoutCache() (string, []byte) {
	if len(*s.Filenames) == 0 {
		s.GetAllFilesFromDir()
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(*s.Filenames))
	key := (*s.Filenames)[randomIndex]

	file, err := os.Open(key)
	if err != nil {
		s.logger.Error("failed to open image file: %v", err.Error())
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		s.logger.Error("failed to read image bytes: %v", err.Error())
	}

	return key, imageBytes
}

func (s *ImageService) RefreshCache() {
	for range imagesCachedPerJob {
		key, imageBytes := s.GetRandomImageWithoutCache()
		s.Redis.SetBytes(key, imageBytes)
	}

}
