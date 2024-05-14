package tiktok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTiktokService_Link(t *testing.T) {
	url := "url"
	service := NewTiktokService(url)
	l := service.Link()
	assert.Equal(t, tiktokServiceName, l.Name)
	assert.Equal(t, url, l.Link)
}
