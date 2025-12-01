package main

import (
	// "encoding/json"
	"log"
	"net/http"

	// "os"
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func getRandom() []byte {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()
	key, err := rdb.RandomKey(ctx).Result()
	if err != nil {
		log.Fatal("")
		panic(err)
	}
	val, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		panic(err)
	}
	return val
}

func getRandomImage(c *gin.Context) {
	data := getRandom()
	c.Data(http.StatusOK, "image/jpeg", data)
}

func getAllFilesFromDir() []string {
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
		return nil
	}

	return output
}

func main() {
	router := gin.Default()
	router.GET("/random", getRandomImage)

	router.Run("localhost:8080")
	// http.HandleFunc("/random", getRandomImage)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
