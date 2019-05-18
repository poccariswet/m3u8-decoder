package m3u8

// NewMedia parse line has #EXT-X-MEDIA
func NewMedia(line string) (*MediaSeqment, error) {
	item := parseLine(line[len(ExtMedia+":"):])
	/*
		type MediaSeqment struct {
			URI        string
			Type       string
			GroupID    string
			Language   string
			Name       string
			Default    string
			Autoselect string
		}
	*/
	return &MediaSeqment{
		URI:        item["URI"],
		Type:       item["TYPE"],
		GroupID:    item["GROUP-ID"],
		Language:   item["LANGUAGE"],
		Name:       item["NAME"],
		Default:    item["DEFAULT"],
		Autoselect: item["AUTOSELECT"],
	}
}

func (ms *MediaSeqment) String() string {
	return "MediaSeqment"
}
