package m3u8_test

import (
	"strings"
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	line := `#EXT-X-KEY:METHOD=AES-256,IV=00000000000,URI="http://example.com/keyfile",
KEYFORMAT="identity",KEYFORMATVERSIONS="1/2/5"`

	key, err := m3u8.NewKey(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "AES-256", key.Method)
	assert.Equal(t, "00000000000", key.IV)
	assert.Equal(t, "http://example.com/keyfile", key.URI)
	assert.Equal(t, "identity", key.KeyFormat)
	assert.Equal(t, "1/2/5", key.KeyFormatVersions)
	assert.Equal(t, strings.Replace(line, "\n", "", 1), key.String())
}
