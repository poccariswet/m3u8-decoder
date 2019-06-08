package m3u8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func NewByteRange(line string) (*ByteRangeSegment, error) {
	/*
		type ByteRangeSegment struct {
			Length  int64
			Offset  int64
			Extflag bool
		}
	*/
	if len(line) == 0 {
		return nil, nil
	}

	if strings.HasPrefix(line, ExtByteRange) {
		line = line[len(ExtByteRange+":"):]
	}

	vals := strings.Split(line, "@")
	if len(vals) != 2 {
		return nil, errors.New("ByteRange value is invalid")
	}

	length, err := strconv.ParseInt(vals[0], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "ParseInt err")
	}

	offset, err := strconv.ParseInt(vals[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "ParseInt err")
	}

	return &ByteRangeSegment{
		Length: length,
		Offset: offset,
	}, nil
}

func (bs *ByteRangeSegment) String() string {
	if bs.Extflag {
		return fmt.Sprintf("%s:%d@%d", ExtByteRange, bs.Length, bs.Offset)
	}
	return fmt.Sprintf("%d@%d", bs.Length, bs.Offset)
}
