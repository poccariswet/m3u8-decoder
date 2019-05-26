package m3u8

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func NewExtInf(line string) (*InfSegment, error) {
	/*
		type InfSegment struct {
			Duration float64
			URI      string
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
	return "InfSegment"
}
