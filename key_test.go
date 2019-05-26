package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	line := `#EXT-X-KEY:METHOD=AES-256,URI="http://example.com/keyfile",
IV=00000000000,KEYFORMAT="identity,KEYFORMATVERSIONS="1/2/5"`

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
}
