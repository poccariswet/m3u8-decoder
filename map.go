package m3u8

import (
	"fmt"
	"strings"

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

	br, err := NewByteRange(item[BYTERANGE])
	if err != nil {
		return nil, errors.Wrap(err, "new byte range")
	}

	return &MapSegment{
		URI:       item[URI],
		ByteRange: br,
	}, nil
}

func (ms *MapSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf(`%s="%s"`, URI, ms.URI))

	if ms.ByteRange != nil {
		s = append(s, fmt.Sprintf(`%s="%s"`, BYTERANGE, ms.ByteRange.String()))
	}
	return fmt.Sprintf("%s:%s", ExtMap, strings.Join(s, ","))
}
