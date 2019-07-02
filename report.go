package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type RenditionReportSegment struct {
	URI      string
	LastMSN  uint64
	LastPART uint64
}

func NewReport(line string) (*RenditionReportSegment, error) {
	item := parseLine(line[len(ExtRenditionReport+":"):])

	lastMSN, err := extractUint64(item, LASTMSN)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}

	lastPART, err := extractUint64(item, LASTPART)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}
	return &RenditionReportSegment{
		URI:      item[URI],
		LastMSN:  lastMSN,
		LastPART: lastPART,
	}, nil
}

func (rs *RenditionReportSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf(`%s="%s"`, URI, rs.URI))

	if rs.LastMSN != 0 {
		s = append(s, fmt.Sprintf("%s=%d", LASTMSN, rs.LastMSN))
	}

	if rs.LastPART != 0 {
		s = append(s, fmt.Sprintf("%s=%d", LASTPART, rs.LastPART))
	}
	return fmt.Sprintf("%s:%s", ExtRenditionReport, strings.Join(s, ","))
}
