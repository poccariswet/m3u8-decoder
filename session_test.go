package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestSessionDate(t *testing.T) {
	line := `#EXT-X-SESSION-DATA:DATA-ID="com.example.title",LANGUAGE="es",VALUE="Este es un ejemplo", URI="http://sample.com"`

	sd, err := m3u8.NewSessionData(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "com.example.title", sd.DataID)
	assert.Equal(t, "Este es un ejemplo", sd.Value)
	assert.Equal(t, "http://sample.com", sd.URI)
	assert.Equal(t, "es", sd.Language)
}

func TestSessionKey(t *testing.T) {
	line := `#EXT-X-SESSION-KEY:METHOD=AES-256,URI="http://example.com/keyfile",
IV=FFFFFF0,KEYFORMAT="identity,KEYFORMATVERSIONS="1/2/5"`

	sk, err := m3u8.NewSessionKey(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "AES-256", sk.Method)
	assert.Equal(t, "FFFFFF0", sk.IV)
	assert.Equal(t, "http://example.com/keyfile", sk.URI)
	assert.Equal(t, "identity", sk.KeyFormat)
	assert.Equal(t, "1/2/5", sk.KeyFormatVersions)
}
