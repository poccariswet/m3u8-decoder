package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// ServerControl indicate support for features such as Blocking Playlist Reload and Playlist Delta Updates
type ServerControlSegment struct {
	CanBlockReload bool
	CanSkipUntil   float64
	HoldBack       float64
	PartHoldBack   float64
}

// New server control segment
func NewServerControl(line string) (*ServerControlSegment, error) {
	item := parseLine(line[len(ExtServerControl+":"):])

	canblock, err := extractBool(item, CANBLOCKRELOAD)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	until, err := extractFloat64(item, CANSKIPUNTIL)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	holdback, err := extractFloat64(item, HOLDBACK)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	parthb, err := extractFloat64(item, PARTHOLDBACK)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	return &ServerControlSegment{
		CanBlockReload: canblock,
		CanSkipUntil:   until,
		HoldBack:       holdback,
		PartHoldBack:   parthb,
	}, nil
}

// segment to string
func (ss *ServerControlSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=YES", CANBLOCKRELOAD))

	if ss.PartHoldBack != 0 {
		s = append(s, fmt.Sprintf("%s=%v", PARTHOLDBACK, ss.PartHoldBack))
	}

	if ss.CanSkipUntil != 0 {
		s = append(s, fmt.Sprintf("%s=%v", CANSKIPUNTIL, ss.CanSkipUntil))
	}

	if ss.HoldBack != 0 {
		s = append(s, fmt.Sprintf("%s=%v", HOLDBACK, ss.HoldBack))
	}

	return fmt.Sprintf("%s:%s", ExtServerControl, strings.Join(s, ","))
}
