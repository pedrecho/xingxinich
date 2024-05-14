package drive

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_UploadFile(t *testing.T) {
	service, err := NewService(context.Background(), "./../../configs/xingxinichbot-de2e63130db2.json")
	assert.NoError(t, err)
	url, err := service.UploadFile("./../../testdata/videoplayback.mp4")
	assert.NoError(t, err)
	fmt.Println(url)
}
