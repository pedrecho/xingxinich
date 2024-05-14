package bot

import (
	tb "gopkg.in/telebot.v3"
	"log"
)

type ShortStep = int

const (
	ShortDefault ShortStep = iota
	ShortTitle             // 0
	ShortDescription
	ShortVideo
)

type ShortState struct {
	Step  ShortStep
	Title string
	Desc  string
}

type ShortMap = map[UserID]*ShortState

func (b *Bot) Short() {
	b.bot.Handle("/short", func(c tb.Context) error {
		userID := c.Sender().ID
		b.shortMap[userID] = &ShortState{
			Step: ShortTitle,
		}
		return c.Send("Please enter the video title:")
	})

	b.bot.Handle(tb.OnText, func(c tb.Context) error {
		userID := c.Sender().ID
		short, ok := b.shortMap[userID]
		if !ok || short.Step == ShortDefault {
			return nil
		}

		switch short.Step {
		case ShortTitle:
			short.Title = c.Text()
			short.Step = ShortDescription
			return c.Send("Now, enter the video description:")
		case ShortDescription:
			short.Desc = c.Text()
			short.Step = ShortVideo
			return c.Send("Finally, send me the video file:")
		default:
			return nil
		}
	})

	b.bot.Handle(tb.OnVideo, func(c tb.Context) error {
		userID := c.Sender().ID
		short, ok := b.shortMap[userID]
		if !ok || short.Step != ShortVideo {
			return nil
		}

		// TODO Process the video here (save it, log it, etc.)
		video := c.Message().Video
		log.Printf("Received video from %d: %s, %s, FileID: %s\n", userID, short.Title, short.Desc, video.FileID)

		delete(b.shortMap, userID)
		return c.Send("Video received and processed!")
	})
}
