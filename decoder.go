package m3u8

import (
	"bytes"
	"io"
)

type PlayList struct{}

type PlayListType uint

// Decode From returns playlist and find the type.
func DecodeFrom(r io.Reader) (PlayList, PlayListType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, 0, err
	}

	return decode(buf)
}
