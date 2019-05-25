package m3u8

import (
	"github.com/pkg/errors"
)

// NewVariant parse line has EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF
func NewVariant(line string) (*VariantSeqment, error) {
	item := parseLine(line)
	/*
		type VariantSeqment struct {
			IFrame bool

			URI              string
			Bandwidth        uint32
			Name             string
			Subtitle         string
			AverageBandwidth uint32
			ProgramID        uint32
			Codec            string
			AudioCodec       string
			Audio            string
			Video            string
			FrameRate        float64
			ClosedCaptions   string
			HDCPLevel        string
			Resolution       *Resolution
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

	return &VariantSeqment{
		URI:              item[URI],
		Bandwidth:        uint32(bandwidth),
		Name:             item[NAME],
		Subtitle:         item[SUBTITLES],
		AverageBandwidth: uint32(averageBandwidth),
		ProgramID:        uint32(programID),
		Codec:            item[CODECS],
		Audio:            item[AUDIO],
		Video:            item[VIDEO],
		FrameRate:        frameRate,
		ClosedCaptions:   item[CLOSEDCAPTIONS],
		HDCPLevel:        item[HDCPLEVEL],
		Resolution:       resolution,
	}, nil
}

func (va *VariantSeqment) String() string {
	return "VariantSeqment"
}
