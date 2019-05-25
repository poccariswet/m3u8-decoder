package m3u8

import (
	"github.com/pkg/errors"
)

func NewMap(line string) (*MapSeqment, error) {
	/*
		type MapSeqment struct {
			URI       string
			ByteRange *ByteRangeSegment
		}
	*/

	item := parseLine(line[len(ExtMap+":"):])

	br, err := NewByteRange(item[ByteRange])
	if err != nil {
		return nil, errors.Wrap(err, "new byte range")
	}

	return &MapSeqment{
		URI:       item[URI],
		ByteRange: br,
	}, nil
}

func (ms *MapSeqment) String() string {
	return "MapSeqment"
}
