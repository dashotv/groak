package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideos(t *testing.T) {
	m := NewMyAnime()
	urls := m.Read("https://myanime.live/tag/perfect-world/")
	assert.NotEmpty(t, urls, "expected results")
}
