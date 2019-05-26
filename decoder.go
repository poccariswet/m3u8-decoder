package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func decode(buf *bytes.Buffer) (*Playlist, error) {
	playlist := NewPlaylist()
	var end bool
	states := new(States)

	for !end {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			end = true
		} else if err != nil {
			return nil, err
		}

		if len(line) < 1 || line == "\r" {
			continue
		}

		line = strings.TrimSpace(line)
		if err := decodeLine(playlist, line, states); err != nil {
			return playlist, err
		}
	}

	return playlist, nil
}

func DecodeFrom(r io.Reader) (*Playlist, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return decode(buf)
}

func decodeLine(p *Playlist, line string, s *States) error {
	if !s.m3u8 && line != EXTM3U {
		return errors.New("invalid playlist, not exist #EXTM3U")
	}

	switch {
	case line == EXTM3U:
		s.m3u8 = true
	case strings.HasPrefix(line, ExtENDList):
		p.live = false
	case strings.HasPrefix(line, ExtVersion):
		_, err := fmt.Sscanf(line, ExtVersion+":%d", &p.version)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, ExtMedia):
		m, err := NewMedia(line)
		if err != nil {
			return errors.Wrap(err, "new media err")
		}
		p.Segments = append(p.Segments, m)
	case strings.HasPrefix(line, ExtStreamInf):
		p.master = true
		s.frameTag = true
		line = line[len(ExtStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	case strings.HasPrefix(line, ExtFrameStreamInf):
		p.master = true
		s.frameTag = false
		line = line[len(ExtFrameStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		v.IFrame = true
		s.segment = v
		p.Segments = append(p.Segments, v)
	case strings.HasPrefix(line, ExtByteRange):
		br, err := NewByteRange(line)
		if err != nil {
			return errors.Wrap(err, "new byte range err")
		}
		_ = br
	case strings.HasPrefix(line, ExtMap):
		m, err := NewMap(line)
		if err != nil {
			return errors.Wrap(err, "new map err")
		}
		p.Segments = append(p.Segments, m)
	case strings.HasPrefix(line, ExtKey):
		key, err := NewKey(line)
		if err != nil {
			return errors.Wrap(err, "new key err")
		}
		p.Segments = append(p.Segments, key)
	default:
	}
	return nil
}
