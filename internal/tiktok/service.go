package tiktok

import (
	"xingxinich/pkg/service"
)

const tiktokServiceName = "Tiktok"

type TiktokService struct {
	link string
}

func NewTiktokService(link string) *TiktokService {
	return &TiktokService{
		link: link,
	}
}

func (s *TiktokService) Short(shortData service.ShortData) error {
	//TODO real work)
	return nil
}

func (s *TiktokService) Link() service.LinkData {
	return service.LinkData{
		Name: tiktokServiceName,
		Link: s.link,
	}
}
func (s *TiktokService) Shutdown() error {
	return nil
}
