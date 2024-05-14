package app

import (
	"fmt"
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

	l := logic.NewLogic()
	defer l.Shutdown()

	// TODO log middleware
	b, err := bot.NewBot(cfg.TgBot.Token, l)
	if err != nil {
		return fmt.Errorf("tg bot initialization: %w", err)
	}
	b.Start()

	return nil
}
