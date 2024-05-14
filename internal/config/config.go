package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"xingxinich/pkg/zaplogger"
)

type Config struct {
	TgBot     TgBotConfig      `yaml:"tg_bot"`
	Cloud     CloudConfig      `yaml:"cloud"`
	Instagram InstagramConfig  `yaml:"instagram"`
	Logger    zaplogger.Config `yaml:"zaplogger"`
}

func Load(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	cfg := Config{}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return &cfg, nil
}

type TgBotConfig struct {
	Token string `yaml:"token"`
}

type CloudConfig struct {
	CredentialsPath string `yaml:"credentials_path"`
}

type InstagramConfig struct {
}
