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
