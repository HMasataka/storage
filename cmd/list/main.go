package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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

	bucketName := "tmp"

	buckets, err := list(ctx, client, bucketName)
	if err != nil {
		log.Fatalf("failed to list: %v", err)
	}
	fmt.Printf("buckets: %+v\n", buckets)
}

func list(ctx context.Context, client *storage.Client, bucketName string) ([]string, error) {
	var objects []string

	it := client.Bucket(bucketName).Objects(ctx, &storage.Query{})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, attrs.Name)
	}

	return objects, nil
}
