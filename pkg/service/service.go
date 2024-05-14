package service

type MediaService interface {
	Short(ShortData) error
	Link() LinkData
	Shutdown() error
}

type ShortData struct {
	Title       string
	Description string
	Link        string
}

type LinkData struct {
	Name string
	Link string
}
