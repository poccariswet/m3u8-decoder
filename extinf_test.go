package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestExtInf(t *testing.T) {
	line := `#EXTINF:9.009,`

	inf, err := m3u8.NewExtInf(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, float64(9.009), inf.Duration)
}
