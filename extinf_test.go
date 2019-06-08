package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestExtInf(t *testing.T) {
	line := `#EXTINF:9.009,
#EXT-X-BYTERANGE:4500@600
1.ts`

	inf, err := m3u8.NewExtInf(line)
	if err != nil {
		t.Fatal(err)
	}

	br, err := m3u8.NewByteRange(`#EXT-X-BYTERANGE:4500@600`)
	if err != nil {
		t.Fatal(err)
	}

	// process in decode.go
	inf.ByteRange = br
	inf.URI = `1.ts`

	assert.Nil(t, err)
	assert.Equal(t, float64(9.009), inf.Duration)
	assert.Equal(t, line, inf.String())
}
