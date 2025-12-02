package service

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageService struct {
	Redis     RedisService
	Filenames *[]string
}

func NewImageService(r RedisService) *ImageService {
	return &ImageService{
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
	imageRootDirectory := "/home/henry/Desktop/images/"

	err := filepath.Walk(imageRootDirectory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		for _, fileType := range fileTypes {
			if strings.HasSuffix(path, fileType) {
				output = append(output, path)
			}
		}

		if info.IsDir() {
			fmt.Printf("visited dir: %q\n", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", imageRootDirectory, err)
	}
	fmt.Printf("output len %d\n", len(output))

	*s.Filenames = output
}

func (s ImageService) GetRandomImage() []byte {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "WARNING: ", log.Lshortfile)

		warnf = func(s string) {
			logger.Output(2, s)
		}
	)

	key := s.Redis.GetRandomKey()
	val, err := s.Redis.GetBytes(key)
	if err != nil {
		warnf("Failed to retrieve image from cache, grabbing image file directly")
		_, val = s.GetRandomImageWithoutCache()
	}
	return val
}

func (s ImageService) GetRandomImageWithoutCache() (string, []byte) {

	log.Printf("file array len: %d", len(*s.Filenames))
	if len(*s.Filenames) == 0 {
		s.GetAllFilesFromDir()
	}

	log.Printf("file array len: %d", len(*s.Filenames))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(*s.Filenames))
	key := (*s.Filenames)[randomIndex]

	file, err := os.Open(key)
	if err != nil {
		log.Fatalf("failed to open image file: %v", err)
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read image bytes: %v", err)
	}

	return key, imageBytes
}

func (s *ImageService) RefreshCache() {
	for i := 0; i < 2; i++ {
		key, imageBytes := s.GetRandomImageWithoutCache()
		s.Redis.SetBytes(key, imageBytes)
	}

}
