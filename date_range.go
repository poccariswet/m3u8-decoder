package m3u8

import (
	"time"

	"github.com/pkg/errors"
)

func NewDateRange(line string) (*DateRangeSegment, error) {
	/*
		type DateRangeSegment struct {
			ID              string
			Class           string
			StartDate       time.Time
			EndDate         time.Time
			Duration        float64
			PlannedDuration float64
			Scte35Cmd       string
			Scte35Out       string
			Scte35In        string
			EndOnNext       bool
		}
	*/

	item := parseLine(line[len(ExtByteRange+":"):])

	duration, err := extractFloat64(item, DURATION)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	plannedDuration, err := extractFloat64(item, PLANNEDDURATION)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	endOnNext, err := extractBool(item, ENDONNEXT)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	s, ok := item[STARTDATE]
	var startDate time.Time
	if ok {
		startDate, err = parseFullTime(s)
		if err != nil {
			return nil, errors.Wrap(err, "parseFullTime err")
		}
	}

	e, ok := item[ENDDATE]
	var endDate time.Time
	if ok {
		endDate, err = parseFullTime(e)
		if err != nil {
			return nil, errors.Wrap(err, "parseFullTime err")
		}
	}

	return &DateRangeSegment{
		ID:              item[ID],
		Class:           item[CLASS],
		StartDate:       startDate,
		EndDate:         endDate,
		Duration:        duration,
		PlannedDuration: plannedDuration,
		Scte35Cmd:       item[SCTE35CMD],
		Scte35Out:       item[SCTE35OUT],
		Scte35In:        item[SCTE35IN],
		EndOnNext:       endOnNext,
	}, nil
}

func (ds *DateRangeSegment) String() string {
	return "DateRangeSegment"
}
