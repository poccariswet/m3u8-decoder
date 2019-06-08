package m3u8

import (
	"fmt"
	"strings"
)

func NewSessionData(line string) (*SessionDataSegment, error) {
	/*
		type SessionDataSegment struct {
			DataID   string
			Value    string
			URI      string
			Language string
		}
	*/

	item := parseLine(line[len(ExtSessionData+":"):])

	return &SessionDataSegment{
		DataID:   item[DATAID],
		Value:    item[VALUE],
		URI:      item[URI],
		Language: item[LANGUAGE],
	}, nil
}

func (ss *SessionDataSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf(`%s="%s"`, DATAID, ss.DataID))

	if ss.Language != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, LANGUAGE, ss.Language))
	}

	if ss.Value != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, VALUE, ss.Value))
	}

	if ss.URI != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, URI, ss.URI))
	}
	return fmt.Sprintf("%s:%s", ExtSessionData, strings.Join(s, ","))
}

func NewSessionKey(line string) (*SessionKeySegment, error) {
	/*
		type SessionKeySegment struct {
			Method            string
			IV                string // Initialization Vector
			URI               string
			KeyFormat         string
			KeyFormatVersions string
		}
	*/

	item := parseLine(line[len(ExtSessionKey+":"):])

	return &SessionKeySegment{
		Method:            item[METHOD],
		IV:                item[IV],
		URI:               item[URI],
		KeyFormat:         item[KEYFORMAT],
		KeyFormatVersions: item[KEYFORMATVERSIONS],
	}, nil

}

func (ss *SessionKeySegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=%s", METHOD, ss.Method))

	if ss.IV != "" {
		s = append(s, fmt.Sprintf("%s=%s", IV, ss.IV))
	}

	if ss.KeyFormat != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, KEYFORMAT, ss.KeyFormat))
	}

	if ss.KeyFormatVersions != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, KEYFORMATVERSIONS, ss.KeyFormatVersions))
	}

	if ss.URI != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, URI, ss.URI))
	}

	return fmt.Sprintf("%s:%s", ExtSessionKey, strings.Join(s, ","))
}
