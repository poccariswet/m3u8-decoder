package m3u8

// NewVariant parse line has EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF
func NewVariant(line string) (*VariantAttributes, error) {
	item := parseLine(line[len(ExtFrameStreamInf+":"):])
	_ = item
	/*
		type VariantAttributes struct {
			Bandwidth  uint64
			ProgramID  uint64
			Codec      string
			Resolution string
			Audio      string
			Video      string
		}
	*/
	return &VariantAttributes{}
}

func (va *VariantAttributes) String() string {
	return "VariantAttributes"
}
