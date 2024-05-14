package logic

import (
	"context"
	"fmt"
	"xingxinich/internal/cloud"
	"xingxinich/internal/config"
)

type Logic struct {
	cloudService *cloud.Service
}

func NewLogic(cfg *config.Config) (*Logic, error) {
	//TODO ctx
	cloudService, err := cloud.NewService(context.Background(), cfg.Cloud.CredentialsPath)
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
