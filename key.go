package m3u8

import (
	"fmt"
	"strings"
)

func NewKey(line string) (*KeySegment, error) {
	/*
		type KeySegment struct {
			Method            string
			IV                string // Initialization Vector
			URI               string
			KeyFormat         string
			KeyFormatVersions string
		}
	*/

	item := parseLine(line[len(ExtKey+":"):])

	return &KeySegment{
		Method:            item[METHOD],
		IV:                item[IV],
		URI:               item[URI],
		KeyFormat:         item[KEYFORMAT],
		KeyFormatVersions: item[KEYFORMATVERSIONS],
	}, nil
}

func (ks *KeySegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=%s", METHOD, ks.Method))

	if ks.IV != "" {
		s = append(s, fmt.Sprintf("%s=%s", IV, ks.IV))
	}

	if ks.URI != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, URI, ks.URI))
	}

	if ks.KeyFormat != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, KEYFORMAT, ks.KeyFormat))
	}

	if ks.KeyFormatVersions != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, KEYFORMATVERSIONS, ks.KeyFormatVersions))

	}
	return fmt.Sprintf("%s:%s", ExtKey, strings.Join(s, ","))
}
