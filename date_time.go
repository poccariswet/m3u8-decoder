package m3u8

import "github.com/pkg/errors"

func NewProgramDateTime(line string) (*DateTimeSegment, error) {
	/*
		type DateTimeSegment struct {
			Time time.Time
		}
	*/

	line = line[len(ExtProgramDateTime+":"):]

	t, err := parseFullTime(line)
	if err != nil {
		return nil, errors.Wrap(err, "parseTime err")
	}

	return &DateTimeSegment{
		Time: t,
	}, nil
}

func (ds *DateTimeSegment) String() string {
	return "DateTimeSegment"
}
