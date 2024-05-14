package app

import (
	"fmt"
	"go.uber.org/zap"
	"xingxinich/internal/bot"
	"xingxinich/internal/config"
	"xingxinich/internal/logic"
	"xingxinich/pkg/zaplogger"
)

func Run(configPath string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("config initialization: %w", err)
	}
	zapsync, err := zaplogger.ReplaceZap(cfg.Logger)
	if err != nil {
		return fmt.Errorf("zaplogger initialization: %w", err)
	}
	defer zapsync()

	l, err := logic.NewLogic(cfg)
	if err != nil {
		return fmt.Errorf("logic initialization: %w", err)
	}
	defer l.Shutdown()

	//TODO log middleware
	b, err := bot.NewBot(cfg.TgBot.Token, l)
	if err != nil {
		return fmt.Errorf("tg bot initialization: %w", err)
	}

	zap.S().Info("Starting bot...")
	b.Start()

	return nil
}
