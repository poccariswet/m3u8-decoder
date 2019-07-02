package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestPartInf(t *testing.T) {
	line := `#EXT-X-PART-INF:PART-TARGET=0.33334`

	pi, err := m3u8.NewPartInf(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, float64(0.33334), pi.PartTartget)
	assert.Equal(t, line, pi.String())
}

func TestPart(t *testing.T) {
	line := `#EXT-X-PART:DURATION=0.33334,URI="filePart272.f.ts",INDEPENDENT=YES,GAP=YES`

	part, err := m3u8.NewPart(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, float64(0.33334), part.Duration)
	assert.Equal(t, "filePart272.f.ts", part.URI)
	assert.True(t, part.Independent)
	assert.Nil(t, part.ByteRange)
	assert.True(t, part.Gap)
	assert.Equal(t, line, part.String())
}
