package m3u8

import "github.com/pkg/errors"

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
	return "StartSegment"
}
