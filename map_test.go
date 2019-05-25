package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	line := `#EXT-X-MAP:URI="main.mp4",BYTERANGE="560@0"`
	m, err := m3u8.NewMap(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "main.mp4", m.URI)
	assert.NotNil(t, m.ByteRange)

	line = `#EXT-X-MAP:URI="main.mp4"`
	m, err = m3u8.NewMap(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "main.mp4", m.URI)
	assert.Nil(t, m.ByteRange)
}
