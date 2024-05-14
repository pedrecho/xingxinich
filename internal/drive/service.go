package drive

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"os"
)

const publicUrl = "https://drive.google.com/uc?export=view&id="

type Service struct {
	//TODO move to pkg?
	service *drive.Service
}

func NewService(ctx context.Context, credentialsPath string) (*Service, error) {
	service, err := drive.NewService(ctx, option.WithCredentialsFile(credentialsPath), option.WithScopes(drive.DriveScope))
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %v", err)
	}

	return &Service{
		service: service,
	}, nil
}

func (s *Service) UploadFile(filename string) (string, error) {
	//TODO add errs descriptions
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("file open: %w", err)
	}
	info, err := f.Stat()
	if err != nil {
		return "", fmt.Errorf("file stat: %w", err)
	}
	defer f.Close()

	file := &drive.File{
		Name: info.Name(),
	}
	res, err := s.service.Files.Create(file).Media(f).Do()
	if err != nil {
		return "", fmt.Errorf("create remote file: %w", err)
	}

	perm := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err = s.service.Permissions.Create(res.Id, perm).Do()
	if err != nil {
		return "", fmt.Errorf("create permission: %w", err)
	}

	return publicUrl + res.Id, nil
}

//TODO remove file

func (s *Service) Shutdown() error {
	return nil
}
