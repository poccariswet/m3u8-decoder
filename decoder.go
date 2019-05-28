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
	case strings.HasPrefix(line, EXTINF):
		inf, err := NewExtInf(line)
		if err != nil {
			return errors.Wrap(err, "new extinf err")
		}
		s.master = false
		s.segment = inf
		s.segmentTag = true
	case strings.HasPrefix(line, ExtMedia):
		m, err := NewMedia(line)
		if err != nil {
			return errors.Wrap(err, "new media err")
		}
		p.Segments = append(p.Segments, m)
	case strings.HasPrefix(line, ExtStreamInf):
		p.master = true
		s.segmentTag = true
		line = line[len(ExtStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	case strings.HasPrefix(line, ExtFrameStreamInf):
		p.master = true
		s.segmentTag = false
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
	case strings.HasPrefix(line, ExtProgramDateTime):
		dt, err := NewProgramDateTime(line)
		if err != nil {
			return errors.Wrap(err, "new program date time err")
		}
		p.Segments = append(p.Segments, dt)
	case strings.HasPrefix(line, ExtDateRange):
		dr, err := NewDateRange(line)
		if err != nil {
			return errors.Wrap(err, "new date range err")
		}
		p.Segments = append(p.Segments, dr)

		/* session tags */
	case strings.HasPrefix(line, ExtSessionKey):
		sk, err := NewSessionKey(line)
		if err != nil {
			return errors.Wrap(err, "new session key err")
		}
		p.Segments = append(p.Segments, sk)
	case strings.HasPrefix(line, ExtSessionData):
		sd, err := NewSessionData(line)
		if err != nil {
			return errors.Wrap(err, "new session data err")
		}
		p.Segments = append(p.Segments, sd)

		/* playlist tags */
	case strings.HasPrefix(line, ExtPlaylistType):
		_, err := fmt.Sscanf(line, ExtPlaylistType+":%s", &p.PlaylistType)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, ExtIFramesOnly):
		p.IFrameOnly = true
	case strings.HasPrefix(line, ExtTargetDutation):
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%f", &p.TargetDuration)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, ExtDiscontinuitySequence):
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%d", &p.DiscontinuitySequence)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, ExtAllowCache):
		p.AllowCache = parseBool(line[len(ExtAllowCache+":"):])
	case strings.HasPrefix(line, ExtMediaSequence):
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%d", &p.MediaSequence)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	default:
	}
	return nil
}
