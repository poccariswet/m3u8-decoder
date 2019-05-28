package m3u8

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
	return "SessionDataSegment"
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
	return "SessionKeySegment"
}
