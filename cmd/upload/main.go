package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	err := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	const (
		bucketName = "tmp"
		objectPath = "sample.txt"
	)

	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, f); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
