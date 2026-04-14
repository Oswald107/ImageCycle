package cache

type CacheService interface {
	GetImageName() (string, error)
	StoreImageName(imageName string) error
	CreateNewSet(imageNames []string) error
}
