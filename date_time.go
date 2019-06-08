package m3u8

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

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
	var s string
	if !ds.Time.IsZero() {
		s = fmt.Sprintf("%s:%s", ExtProgramDateTime, ds.Time.Format(time.RFC3339Nano))
	}

	return s
}
