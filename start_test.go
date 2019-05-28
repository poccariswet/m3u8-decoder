package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	line := `#EXT-X-START:TIME-OFFSET=25,PRECISE=YES`

	start, err := m3u8.NewStart(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, float64(25), start.TimeOffset)
	assert.True(t, start.Precise)
}
