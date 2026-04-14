package storage

import (
	"errors"
	"fmt"
	"math/rand"

	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type FileStoreService struct {
	filenames []string
	logger    *slog.Logger
}

func NewFileStoreService(logger *slog.Logger) *FileStoreService {
	// get files
	filenames := []string{}
	fileTypes := []string{
		".jpg",
		".jpeg",
		".png",
		".gif",
		".bmp",
		".tiff",
		".webp",
	}
	imageRootDirectory := "/home/henry/Desktop/images"

	err := filepath.Walk(imageRootDirectory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			logger.Warn("failure accessing a path %q: %v\n", path, err)
			return err
		}

		for _, fileType := range fileTypes {
			if strings.HasSuffix(path, fileType) {
				filenames = append(filenames, path)
			}
		}

		if info.IsDir() {
			logger.Debug("visited dir: %q\n", path)
		}
		return nil
	})
	if err != nil {
		logger.Warn("error walking the path %q: %v\n", imageRootDirectory, err)
	}
	logger.Debug("output len %d\n", len(filenames))

	return &FileStoreService{
		filenames: filenames,
		logger:    logger.With("FileStoreService"),
	}
}

func (fs *FileStoreService) GetImage(imageName string) ([]byte, error) {
	if len(fs.filenames) == 0 {
		// problem
	}

	file, err := os.Open(imageName)
	if err != nil {
		fs.logger.Error("failed to open image file: %v", err.Error())
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		fs.logger.Error("failed to read image bytes: %v", err.Error())
	}

	return imageBytes, nil
}

func (fs *FileStoreService) GetRandomImageName() (string, error) {
	if len(fs.filenames) == 0 {
		return "", errors.New("Map is empty, no random key to select.")
	}

	randomIndex := rand.Intn(len(fs.filenames))
	randomFilename := fs.filenames[randomIndex]

	fmt.Printf("Random Key: %s, \n", randomFilename)
	return randomFilename, nil
}
