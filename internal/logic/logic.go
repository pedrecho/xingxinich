package logic

import (
	"context"
	"fmt"
	"xingxinich/internal/config"
	"xingxinich/internal/drive"
)

type Logic struct {
	cloudService *drive.Service
}

func NewLogic(cfg *config.Config) (*Logic, error) {
	//TODO ctx
	cloudService, err := drive.NewService(context.Background(), cfg.Drive.CredentialsPath)
	if err != nil {
		return nil, fmt.Errorf("google cloud initialization: %w", err)
	}
	return &Logic{
		cloudService: cloudService,
	}, nil
}

func (l *Logic) Shutdown() error {
	return l.cloudService.Shutdown()
}
