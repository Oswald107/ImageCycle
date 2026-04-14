package storage

type StorageService interface {
	GetImage(imageName string) ([]byte, error)
	GetRandomImageName() (string, error)
}
