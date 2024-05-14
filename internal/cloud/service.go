package cloud

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"os"
)

type Service struct {
	//TODO move to pkg?
	client *storage.Client
}

func NewService(ctx context.Context, credentialsPath string) (*Service, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &Service{
		client: client,
	}, nil
}

func (s *Service) UploadFile(ctx context.Context, bucket, object, filePath string) (string, error) {
	//TODO add errs descriptions
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	wc := s.client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return "", err
	}
	if err = wc.Close(); err != nil {
		return "", err
	}

	acl := s.client.Bucket(bucket).Object(object).ACL()
	if err = acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, object), nil
}

//TODO remove file

func (s *Service) Shutdown() error {
	return s.client.Close()
}
