package logic

import "xingxinich/internal/cloud"

type Logic struct {
	cloudService *cloud.Service
}

func NewLogic() *Logic {
	return &Logic{}
}

func (l *Logic) Shutdown() error {
	return l.cloudService.Shutdown()
}
