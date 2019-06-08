package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// NewMedia parse line has #EXT-X-MEDIA
func NewMedia(line string) (*MediaSegment, error) {
	item := parseLine(line[len(ExtMedia+":"):])
	/*
		type MediaSegment struct {
			Type            string
			GroupID         string
			Language        string
			AssocLanguage   string
			Name            string
			Autoselect      bool
			Forced          bool
			Default         bool
			URI             string
			InstreamID      string
			Characteristics string
			Channels        string
		}
	*/

	autoselect, err := extractBool(item, AUTOSELECT)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	forced, err := extractBool(item, FORCED)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	defaul, err := extractBool(item, DEFAULT)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err")
	}

	return &MediaSegment{
		Type:            item[TYPE],
		GroupID:         item[GROUPID],
		Language:        item[LANGUAGE],
		AssocLanguage:   item[ASSOCLANGUAGE],
		Name:            item[NAME],
		Autoselect:      autoselect,
		Forced:          forced,
		Default:         defaul,
		URI:             item[URI],
		InstreamID:      item[INSTREAMID],
		Characteristics: item[CHARACTERISTICS],
		Channels:        item[CHANNELS],
	}, nil
}

func (ms *MediaSegment) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%s=%s", TYPE, ms.Type))

	if ms.GroupID != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, GROUPID, ms.GroupID))
	}

	if ms.Language != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, LANGUAGE, ms.Language))
	}

	if ms.AssocLanguage != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, ASSOCLANGUAGE, ms.AssocLanguage))
	}

	if ms.Name != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, NAME, ms.Name))
	}

	if ms.Autoselect {
		s = append(s, fmt.Sprintf("%s=YES", AUTOSELECT))
	}

	if ms.Forced {
		s = append(s, fmt.Sprintf("%s=YES", FORCED))
	}

	if ms.Default {
		s = append(s, fmt.Sprintf("%s=YES", DEFAULT))
	}

	if ms.URI != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, URI, ms.URI))
	}

	if ms.InstreamID != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, INSTREAMID, ms.InstreamID))
	}

	if ms.Characteristics != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, CHARACTERISTICS, ms.Characteristics))
	}

	if ms.Channels != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, CHANNELS, ms.Channels))
	}

	return fmt.Sprintf("%s:%s", ExtMedia, strings.Join(s, ","))
}
