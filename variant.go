package m3u8

import "github.com/pkg/errors"

// NewVariant parse line has EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF
func NewVariant(line string) (*VariantAttributes, error) {
	item := parseLine(line[len(ExtFrameStreamInf+":"):])
	/*
			type VariantAttributes struct {
				URI              string
				Name             string
				Subtitle         string
				Bandwidth        uint32
				AverageBandwidth uint32
				ProgramID        uint32
				Codec            string
				AudioCodec       string
				Audio            string
				Video            string
				FrameRate        float64
				HDCPLevel        string
				Resolution       *Resolution

				IFrame bool
		}
	*/

	resolution, err := NewResolution(item, RESOLUTION)
	if err != nil {
		return nil, errors.Wrap(err, "resolution parse err")
	}

	bandwidth, err := extractUint64(item, BANDWIDTH)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}

	averageBandwidth, err := extractUint64(item, AVERAGEBANDWIDTH)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}

	programID, err := extractUint64(item, PROGRAMID)
	if err != nil {
		return nil, errors.Wrap(err, "extractUint64 err")
	}

	frameRate, err := extractFloat64(item, FRAMERATE)
	if err != nil {
		return nil, errors.Wrap(err, "extractFloat64 err")
	}

	return &VariantAttributes{
		URI:              item[URI],
		Name:             item[NAME],
		Subtitle:         item[SUBTITLES],
		Bandwidth:        uint32(bandwidth),
		AverageBandwidth: uint32(averageBandwidth),
		ProgramID:        uint32(programID),
		Codec:            item[CODECS],
		Audio:            item[AUDIO],
		Video:            item[VIDEO],
		FrameRate:        frameRate,
		HDCPLevel:        item[HDCPLEVEL],
		Resolution:       resolution,
	}, nil
}

func (va *VariantAttributes) String() string {
	return "VariantAttributes"
}
