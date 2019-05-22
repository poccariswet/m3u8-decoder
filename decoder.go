package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func decode(buf *bytes.Buffer) (*Playlist, ListType, error) {
	playlist := NewPlaylist()
	var end bool
	states := new(States)

	for !end {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			end = true
		} else if err != nil {
			return nil, ERRTYPE, err
		}

		if len(line) < 1 || line == "\r" {
			continue
		}

		line = strings.TrimSpace(line)
		if err := decodeLine(playlist, line, states); err != nil {
			return playlist, ERRTYPE, err
		}
	}

	switch states.listtype {
	case MASTER:
		return playlist, MASTER, nil
	case MEDIA:
		return playlist, MEDIA, nil
	}

	return nil, ERRTYPE, errors.New("not playlist")
}

func DecodeFrom(r io.Reader) (*Playlist, ListType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, ERRTYPE, err
	}
	return decode(buf)
}

func decodeLine(p *Playlist, line string, s *States) error {
	if s.m3u8 && line != EXTM3U {
		return errors.New("invalid playlist, not exist #EXTM3U")
	}
	switch {
	case line == EXTM3U:
		s.m3u8 = true
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
		s.master = true
		s.frameTag = true
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	case strings.HasPrefix(line, ExtFrameStreamInf):
		s.master = true
		s.frameTag = false
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		v.IFrame = true
		s.segment = v
		p.Segments = append(p.Segments, v)
	case strings.HasPrefix(line, ExtByteRange):
		v := trimLine(line, ExtByteRange+":")
		br, err := NewByteRange(v)
		if err != nil {
			return errors.Wrap(err, "new byte range err")
		}
		_ = br
	}
	return nil
}
