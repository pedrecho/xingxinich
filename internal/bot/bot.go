package bot

import (
	tb "gopkg.in/telebot.v3"
	"time"
	"xingxinich/internal/logic"
)

type UserID = int64

type Bot struct {
	bot      *tb.Bot
	logic    *logic.Logic
	shortMap ShortMap
}

func NewBot(token string, logic *logic.Logic) (*Bot, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &Bot{
		bot:      bot,
		logic:    logic,
		shortMap: make(ShortMap),
	}, nil
}

func (b *Bot) Start() {
	b.Short()
}
