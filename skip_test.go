package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	line := `#EXT-X-SKIP:SKIPPED-SEGMENTS=3`

	skip, err := m3u8.NewSkip(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint64(3), skip.SkippedSegments)
	assert.Equal(t, line, skip.String())
}
