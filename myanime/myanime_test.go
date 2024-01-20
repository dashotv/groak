package myanime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideos(t *testing.T) {
	m := New("https://myanime.live/tag/perfect-world/")
	urls := m.Read()
	assert.NotEmpty(t, urls, "expected results")
}
