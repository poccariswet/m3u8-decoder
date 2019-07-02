package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestParseResolution(t *testing.T) {
	line := map[string]string{
		"RESOLUTION": "416x234",
	}

	r, err := m3u8.NewResolution(line, m3u8.RESOLUTION)
	if err != nil {
		t.Fatalf("err is not nil: %s", err)
	}
	assert.Equal(t, uint16(416), r.Width)
	assert.Equal(t, uint16(234), r.Height)

	line = map[string]string{
		"RESOLUTION": "41ax234",
	}

	r, err = m3u8.NewResolution(line, m3u8.RESOLUTION)
	assert.Error(t, err)
	assert.Nil(t, r)
}
