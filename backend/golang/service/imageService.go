package service

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type ImageService struct {
	Redis     RedisService
	Filenames []string
}

func NewImageService(r RedisService) *ImageService {
	return &ImageService{
		Redis:     r,
		Filenames: []string{},
	}
}

func (s ImageService) GetAllFilesFromDir() {
	output := []string{}
	fileTypes := []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
		"bmp",
		"tiff",
		"webp",
	}
	imageRootDirectory := "/home/henry/Desktop/images/"

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		for _, prefix := range fileTypes {
			if strings.HasPrefix(path, prefix) {
				output = append(output, path)
			}
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", imageRootDirectory, err)
	}

	s.Filenames = output
}
