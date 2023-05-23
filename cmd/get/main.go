package main

import (
	"context"
	"fmt"
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
	defer client.Close()

	const (
		bucketName = "tmp"
		objectPath = "sample.txt"
	)

	_, err = getMetadata(client, ctx, bucketName, objectPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}

func getMetadata(client *storage.Client, ctx context.Context, bucket, object string) (*storage.ObjectAttrs, error) {
	o := client.Bucket(bucket).Object(object)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).Attrs: %v", object, err)
	}

	fmt.Printf("Bucket: %v\n", attrs.Bucket)
	fmt.Printf("CacheControl: %v\n", attrs.CacheControl)
	fmt.Printf("ContentDisposition: %v\n", attrs.ContentDisposition)
	fmt.Printf("ContentEncoding: %v\n", attrs.ContentEncoding)
	fmt.Printf("ContentLanguage: %v\n", attrs.ContentLanguage)
	fmt.Printf("ContentType: %v\n", attrs.ContentType)
	fmt.Printf("Crc32c: %v\n", attrs.CRC32C)
	fmt.Printf("Generation: %v\n", attrs.Generation)
	fmt.Printf("KmsKeyName: %v\n", attrs.KMSKeyName)
	fmt.Printf("Md5Hash: %v\n", attrs.MD5)
	fmt.Printf("MediaLink: %v\n", attrs.MediaLink)
	fmt.Printf("Metageneration: %v\n", attrs.Metageneration)
	fmt.Printf("Name: %v\n", attrs.Name)
	fmt.Printf("Size: %v\n", attrs.Size)
	fmt.Printf("StorageClass: %v\n", attrs.StorageClass)
	fmt.Printf("TimeCreated: %v\n", attrs.Created)
	fmt.Printf("Updated: %v\n", attrs.Updated)
	fmt.Printf("Event-based hold enabled? %t\n", attrs.EventBasedHold)
	fmt.Printf("Temporary hold enabled? %t\n", attrs.TemporaryHold)
	fmt.Printf("Retention expiration time %v\n", attrs.RetentionExpirationTime)
	fmt.Printf("Custom time %v\n", attrs.CustomTime)
	fmt.Printf("\n\nMetadata\n")
	for key, value := range attrs.Metadata {
		fmt.Printf("\t%v = %v\n", key, value)
	}

	return attrs, nil
}
