package bot

type Short = int

const (
	ShortName = iota // 0
	ShortDescription
	ShortVideo
)

type ShortMap = map[int64]Short

func (b *Bot) Short() {

}
