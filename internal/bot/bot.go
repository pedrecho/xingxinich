package bot

import (
	tb "gopkg.in/telebot.v3"
	"time"
)

type Bot struct {
	bot      *tb.Bot
	shortMap ShortMap
}

func NewBot(token string) (*Bot, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &Bot{
		bot:      bot,
		shortMap: make(ShortMap),
	}, nil
}
