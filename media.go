package m3u8

import "github.com/pkg/errors"

// NewMedia parse line has #EXT-X-MEDIA
func NewMedia(line string) (*MediaSeqment, error) {
	item := parseLine(line[len(ExtMedia+":"):])
	/*
		type MediaSeqment struct {
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
		return nil, errors.Wrap(err, "extractBool err: ")
	}

	forced, err := extractBool(item, FORCED)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err: ")
	}

	defaul, err := extractBool(item, DEFAULT)
	if err != nil {
		return nil, errors.Wrap(err, "extractBool err: ")
	}

	return &MediaSeqment{
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

func (ms *MediaSeqment) String() string {
	return "MediaSeqment"
}
