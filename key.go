package m3u8

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
	return "KeySegment"
}
