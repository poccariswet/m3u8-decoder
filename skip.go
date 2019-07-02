package m3u8

import (
	"fmt"

	"github.com/pkg/errors"
)

type SkipSegment struct {
	SkippedSegments uint64
}

func NewSkip(line string) (*SkipSegment, error) {
	item := parseLine(line[len(ExtSkip+":"):])

	skip, err := extractUint64(item, SKIPPEDSEGMENTS)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}
	return &SkipSegment{
		SkippedSegments: skip,
	}, nil
}

func (ss *SkipSegment) String() string {
	return fmt.Sprintf("%s:%s=%d", ExtSkip, SKIPPEDSEGMENTS, ss.SkippedSegments)
}
