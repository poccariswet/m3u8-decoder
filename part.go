package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type PartInfSegment struct {
	PartTartget float64
}

func NewPartInf(line string) (*PartInfSegment, error) {
	item := parseLine(line[len(ExtPartInf+":"):])

	target, err := extractFloat64(item, PARTTARGET)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	return &PartInfSegment{
		PartTartget: target,
	}, nil
}

func (ps *PartInfSegment) String() string {
	return fmt.Sprintf("%s:%s=%v", ExtPartInf, PARTTARGET, ps.PartTartget)
}

type PartSegment struct {
	Duration    float64
	URI         string
	Independent bool
	ByteRange   *ByteRangeSegment
	Gap         bool
}

func NewPart(line string) (*PartSegment, error) {
	item := parseLine(line[len(ExtPart+":"):])

	duration, err := extractFloat64(item, DURATION)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	independent, err := extractBool(item, INDEPENDENT)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	br, err := NewByteRange(item[BYTERANGE])
	if err != nil {
		return nil, errors.Wrap(err, "new byte range")
	}

	gap, err := extractBool(item, GAP)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	return &PartSegment{
		Duration:    duration,
		URI:         item[URI],
		Independent: independent,
		ByteRange:   br,
		Gap:         gap,
	}, nil
}

func (ps *PartSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=%v", DURATION, ps.Duration))

	s = append(s, fmt.Sprintf(`%s="%s"`, URI, ps.URI))

	if ps.Independent {
		s = append(s, fmt.Sprintf("%s=YES", INDEPENDENT))
	}

	if ps.ByteRange != nil {
		s = append(s, fmt.Sprintf(`%s="%s"`, BYTERANGE, ps.ByteRange.String()))
	}

	if ps.Gap {
		s = append(s, fmt.Sprintf("%s=YES", GAP))
	}

	return fmt.Sprintf("%s:%s", ExtPart, strings.Join(s, ","))
}
