package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func NewStart(line string) (*StartSegment, error) {
	/*
		type StartSegment struct {
			TimeOffset float64
			Precise    bool
		}
	*/

	item := parseLine(line[len(ExtStart+":"):])

	timeOffset, err := extractFloat64(item, TIMEOFFSET)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	precise, err := extractBool(item, PRECISE)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	return &StartSegment{
		TimeOffset: timeOffset,
		Precise:    precise,
	}, nil
}

func (ss *StartSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=%v", TIMEOFFSET, ss.TimeOffset))

	if ss.Precise {
		s = append(s, fmt.Sprintf("%s=YES", PRECISE))
	}

	return fmt.Sprintf("%s:%s", ExtStart, strings.Join(s, ","))
}
