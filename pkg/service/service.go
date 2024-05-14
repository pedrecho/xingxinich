package service

type VideoService interface {
	Short(ShortData) error
	Link() LinkData
	Shutdown() error
}

type ShortData struct {
	Title       string
	Description string
	//TODO reader/writer
	//Video
	//TODO or video link?
	//Link
}

type LinkData struct {
	Name string
	Link string
}
