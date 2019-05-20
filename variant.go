package m3u8

// NewVariant parse line has EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF
func NewVariant(line string) (*VariantAttributes, error) {
	item := parseLine(line[len(ExtFrameStreamInf+":"):])
	/*
		type VariantAttributes struct {
			URI        string
			Name       string
			Subtitle   string
			Bandwidth  uint64
			ProgramID  uint64
			Codec      string
			AudioCodec string
			Resolution string
			Audio      string
			Video      string
			FrameRate  float64
			HDCPLevel  string
			Resolution *Resolution

			IFrame bool
		}
	*/

	resolution, err := NewResolution(item[RESOLUTION])
	if !has {
		resolution = nil
	}

	return &VariantAttributes{
		Resolution: resolution,
	}
}

func (va *VariantAttributes) String() string {
	return "VariantAttributes"
}
