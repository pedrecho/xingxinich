package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	//TODO remove this file
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("./fresh-iridium-423214-n8-941d10d980b8.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Настройки
	bucketName := "xingmonkey"
	objectName := "parasha"
	localFilePath := "videoplayback.mp4"

	// Загрузка файла
	err = uploadFile(ctx, client, bucketName, objectName, localFilePath)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// Получение общедоступной ссылки
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	fmt.Println("Public URL:", url)
}

func uploadFile(ctx context.Context, client *storage.Client, bucket, object, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	// Сделать файл общедоступным
	acl := client.Bucket(bucket).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil
}
