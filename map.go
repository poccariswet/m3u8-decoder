package m3u8

import (
	"github.com/pkg/errors"
)

func NewMap(line string) (*MapSegment, error) {
	/*
		type MapSegment struct {
			URI       string
			ByteRange *ByteRangeSegment
		}
	*/

	item := parseLine(line[len(ExtMap+":"):])

	br, err := NewByteRange(item[ByteRange])
	if err != nil {
		return nil, errors.Wrap(err, "new byte range")
	}

	return &MapSegment{
		URI:       item[URI],
		ByteRange: br,
	}, nil
}

func (ms *MapSegment) String() string {
	return "MapSegment"
}
