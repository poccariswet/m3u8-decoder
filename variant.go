package m3u8

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// NewVariant parse line has EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF
func NewVariant(line string) (*VariantSegment, error) {
	item := parseLine(line)
	/*
		type VariantSegment struct {
			IFrame bool

			URI              string
			Bandwidth        uint32
			Name             string
			Subtitle         string
			AverageBandwidth uint32
			ProgramID        uint32
			Codecs           string
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

	return &VariantSegment{
		URI:              item[URI],
		Bandwidth:        uint32(bandwidth),
		Name:             item[NAME],
		Subtitle:         item[SUBTITLES],
		AverageBandwidth: uint32(averageBandwidth),
		ProgramID:        uint32(programID),
		Codecs:           item[CODECS],
		Audio:            item[AUDIO],
		Video:            item[VIDEO],
		FrameRate:        frameRate,
		ClosedCaptions:   item[CLOSEDCAPTIONS],
		HDCPLevel:        item[HDCPLEVEL],
		Resolution:       resolution,
	}, nil
}

func (vs *VariantSegment) String() string {
	var s []string

	if vs.AverageBandwidth != 0 {
		s = append(s, fmt.Sprintf("%s=%d", AVERAGEBANDWIDTH, vs.AverageBandwidth))
	}

	if vs.Bandwidth != 0 {
		s = append(s, fmt.Sprintf("%s=%d", BANDWIDTH, vs.Bandwidth))
	}

	if vs.Audio != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, AUDIO, vs.Audio))
	}

	if vs.Video != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, VIDEO, vs.Video))
	}

	if vs.Codecs != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, CODECS, vs.Codecs))
	}

	if vs.Resolution != nil {
		s = append(s, vs.Resolution.String())
	}

	if vs.FrameRate != 0 {
		s = append(s, fmt.Sprintf("%s=%v", FRAMERATE, vs.FrameRate))
	}

	if vs.ClosedCaptions != "" {
		s = append(s, fmt.Sprintf("%s=%s", CLOSEDCAPTIONS, vs.ClosedCaptions))
	}

	if vs.ProgramID != 0 {
		s = append(s, fmt.Sprintf("%s=%d", PROGRAMID, vs.ProgramID))
	}

	if vs.Name != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, NAME, vs.Name))
	}

	if vs.HDCPLevel != "" {
		s = append(s, fmt.Sprintf("%s=%s", HDCPLEVEL, vs.HDCPLevel))
	}

	if vs.Subtitle != "" {
		s = append(s, fmt.Sprintf(`%s="%s"`, SUBTITLES, vs.Subtitle))
	}

	if vs.IFrame {
		if vs.URI != "" {
			s = append(s, fmt.Sprintf(`%s="%s"`, URI, vs.URI))
		}

		return fmt.Sprintf("%s:%s", ExtFrameStreamInf, strings.Join(s, ","))
	}
	return fmt.Sprintf("%s:%s\n%s", ExtStreamInf, strings.Join(s, ","), vs.URI)
}
