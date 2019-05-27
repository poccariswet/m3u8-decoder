package m3u8_test

import (
	"testing"
	"time"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestDateTime(t *testing.T) {
	line := `#EXT-X-PROGRAM-DATE-TIME:2017-06-09T04:59:01.797Z`

	dt, err := m3u8.NewProgramDateTime(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)

	expected, err := time.Parse("2006-01-02T15:04:05.999999999Z07", "2017-06-09T04:59:01.797Z")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, dt.Time)
}
