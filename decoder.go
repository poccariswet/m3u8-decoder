package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// decode parses a playlist
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

// DecodeFrom read a playlist passed from the io.Reader
func DecodeFrom(r io.Reader) (*Playlist, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return decode(buf)
}

// ReadFile reads contents from filepath and return Playlist
func ReadFile(path string) (*Playlist, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "ReadFile err")
	}

	return decode(bytes.NewBuffer(file))
}

// decodeLine decodes a line of playlist and parses
func decodeLine(p *Playlist, line string, s *States) error {
	if !s.m3u8 && !EXTM3U.match(line) {
		return errors.New("invalid playlist, not exist #EXTM3U")
	}

	switch {
	case EXTM3U.match(line):
		s.m3u8 = true
	case EXTENDLIST.match(line):
		p.live = false
	case EXTVERSION.match(line):
		p.hasVersion = true
		_, err := fmt.Sscanf(line, "#EXT-X-VERSION:%d", &p.Version)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case EXTINF.match(line):
		inf, err := NewExtInf(line)
		if err != nil {
			return errors.Wrap(err, "new extinf err")
		}
		p.master = false
		s.segment = inf
		s.segmentTag = true
	case ExtMedia.match(line):
		m, err := NewMedia(line)
		if err != nil {
			return errors.Wrap(err, "new media err")
		}
		p.AppendSegment(m)
	case ExtStreamInf.match(line):
		p.master = true
		s.segmentTag = true
		line = line[len(ExtStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	case ExtFrameStreamInf.match(line):
		p.master = true
		s.segmentTag = false
		line = line[len(ExtFrameStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		v.IFrame = true
		s.segment = v
		p.AppendSegment(v)
	case ExtByteRange.match(line):
		br, err := NewByteRange(line)
		if err != nil {
			return errors.Wrap(err, "new byte range err")
		}
		br.Extflag = true
		if m, has := s.segment.(*MapSegment); has {
			m.ByteRange = br
			s.segment = m
			br.Extflag = false
		} else if inf, has := s.segment.(*InfSegment); has {
			inf.ByteRange = br
			s.segment = inf
		}
	case ExtMap.match(line):
		m, err := NewMap(line)
		if err != nil {
			return errors.Wrap(err, "new map err")
		}
		p.AppendSegment(m)
	case ExtKey.match(line):
		key, err := NewKey(line)
		if err != nil {
			return errors.Wrap(err, "new key err")
		}
		p.AppendSegment(key)
	case ExtProgramDateTime.match(line):
		dt, err := NewProgramDateTime(line)
		if err != nil {
			return errors.Wrap(err, "new program date time err")
		}
		p.AppendSegment(dt)
	case ExtDateRange.match(line):
		dr, err := NewDateRange(line)
		if err != nil {
			return errors.Wrap(err, "new date range err")
		}
		p.AppendSegment(dr)

	/* low-latency tags */
	case ExtServerControl.match(line):
		sc, err := NewServerControl(line)
		if err != nil {
			return errors.Wrap(err, "new server control err")
		}
		p.AppendSegment(sc)
	case ExtPartInf.match(line):
		pi, err := NewPartInf(line)
		if err != nil {
			return errors.Wrap(err, "new part inf err")
		}
		p.AppendSegment(pi)
	case ExtRenditionReport.match(line):
		report, err := NewReport(line)
		if err != nil {
			return errors.Wrap(err, "new rendition report err")
		}
		p.AppendSegment(report)
	case ExtSkip.match(line):
		skip, err := NewSkip(line)
		if err != nil {
			return errors.Wrap(err, "new skip err")
		}
		p.AppendSegment(skip)
	case ExtPart.match(line):
		part, err := NewPart(line)
		if err != nil {
			return errors.Wrap(err, "new part err")
		}
		p.AppendSegment(part)

	/* session tags */
	case ExtSessionKey.match(line):
		sk, err := NewSessionKey(line)
		if err != nil {
			return errors.Wrap(err, "new session key err")
		}
		p.AppendSegment(sk)
	case ExtSessionData.match(line):
		sd, err := NewSessionData(line)
		if err != nil {
			return errors.Wrap(err, "new session data err")
		}
		p.AppendSegment(sd)

	case ExtStart.match(line):
		start, err := NewStart(line)
		if err != nil {
			return errors.Wrap(err, "new start err")
		}
		p.AppendSegment(start)
	case ExtIndependentSegments.match(line):
		p.IndependentSegments = true

		/* playlist tags */
	case ExtPlaylistType.match(line):
		if err := p.scanLineValue(line, ExtPlaylistType); err != nil {
			return errors.Wrap(err, "invalid playlist type")
		}
	case ExtIFramesOnly.match(line):
		p.IFrameOnly = true
	case ExtTargetDutation.match(line):
		if err := p.scanLineValue(line, ExtTargetDutation); err != nil {
			return errors.Wrap(err, "invalid target duration")
		}
	case ExtDiscontinuitySequence.match(line):
		if err := p.scanLineValue(line, ExtDiscontinuitySequence); err != nil {
			return errors.Wrap(err, "invalid discontinuity sequence")
		}
	case ExtAllowCache.match(line):
		p.AllowCache = parseBool(line[len(ExtAllowCache+":"):])
	case ExtMediaSequence.match(line):
		if err := p.scanLineValue(line, ExtMediaSequence); err != nil {
			return errors.Wrap(err, "invalid media sequence")
		}
	default:
		line = strings.Trim(line, "\n")
		uri := strings.TrimSpace(line)
		if s.segment != nil && s.segmentTag {
			if p.master {
				v, has := s.segment.(*VariantSegment)
				if !has {
					return errors.New("invalid variant playlist")
				}
				v.URI = uri
				p.AppendSegment(v)
			} else {
				i, has := s.segment.(*InfSegment)
				if !has {
					return errors.New("invalid EXTINF segment")
				}
				i.URI = uri
				p.AppendSegment(i)
			}
			s.segmentTag = false

			return nil
		}
	}
	return nil
}
