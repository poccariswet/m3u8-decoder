package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestByteRange(t *testing.T) {
	line := `#EXT-X-BYTERANGE:139684@138744`

	br, err := m3u8.NewByteRange(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, int64(139684), br.Length)
	assert.Equal(t, int64(138744), br.Offset)

	line = `#EXT-X-BYTERANGE:139684@1387a44`
	br, err = m3u8.NewByteRange(line)
	assert.Error(t, err, `ParseInt err: strconv.ParseInt: parsing "1387a44": invalid syntax`)

	line = `#EXT-X-BYTERANGE:139684=138744`
	br, err = m3u8.NewByteRange(line)
	assert.Error(t, err, "ByteRange value is invalid")
}
