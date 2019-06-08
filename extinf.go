package m3u8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func NewExtInf(line string) (*InfSegment, error) {
	/*
		type InfSegment struct {
			Duration  float64
			URI       string
			ByteRange *ByteRangeSegment
		}
	*/
	val := line[len(EXTINF+":"):]
	v := strings.Split(val, ",")

	duration, err := strconv.ParseFloat(v[0], 64)
	if err != nil {
		return nil, errors.Wrap(err, "ParseFloat64 err")
	}

	return &InfSegment{
		Duration: duration,
	}, nil
}

func (is *InfSegment) String() string {
	var b string
	if is.ByteRange != nil {
		b = fmt.Sprintf("\n%s:%s", ExtByteRange, is.ByteRange.String())
	}
	return fmt.Sprintf("%s:%v,%s\n%s", EXTINF, is.Duration, b, is.URI)
}
