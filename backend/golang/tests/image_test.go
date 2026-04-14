package test

import (
	"my-go-api/factory"
	"my-go-api/mocks"

	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSaveUser(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the Database interface
	mockCache := mocks.NewMockCacheService(ctrl)

	// Set expectations on the mock
	mockCache.EXPECT().GetImageName().Return("file1", nil).Times(1)

	// Create a mock of the Database interface
	mockStorage := mocks.NewMockStorageService(ctrl)

	expected := []byte("hello world")
	// Set expectations on the mock
	mockStorage.EXPECT().GetImage("file1").Return(expected, nil).Times(1)

	// Create an instance of your struct and inject the mock
	imageFactory := factory.NewImageFactory(mockStorage, mockCache)

	// Call the method you're testing
	imageName := imageFactory.GetRandomImageName()
	if imageName != "file1" {
		t.Fatalf("failed to return image")
	}
	imageContent := imageFactory.GetImage(imageName)
	if !bytes.Equal(imageContent, expected) {
		t.Fatalf("got %q, want %q", imageContent, expected)
	}
}
