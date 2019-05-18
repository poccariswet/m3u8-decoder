package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func decode(buf *bytes.Buffer) (Playlist, ListType, error) {
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

func DecodeFrom(r io.Reader) (Playlist, ListType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, ERRTYPE, err
	}
	return decode(buf)
}

func decodeLine(p *Playlist, line string, s *States) error {
	if p.m3u && line != EXTM3U {
		return errors.New("invalid playlist, not exist #EXTM3U")
	}
	switch {
	case line == EXTM3U:
		p.m3u = true
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
		s.segmentTag = true
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	}
}

//func (p *MasterPlaylist) decodePlaylist(states *States, line string) error {
//	switch {
//	case line == "#EXTM3U":
//		states.m3u = true
//	case strings.HasPrefix(line, "#EXT-X-MEDIA:"):
//		states.listtype = MASTER
//		for i, v := range parseLine(line[len("#EXT-X-MEDIA:"):]) {
//			switch i {
//			/*
//				type XMedia struct {
//					URI        string
//					Type       string
//					GroupID    string
//					Language   string
//					Name       string
//					Default    string
//					Autoselect string
//				}
//			*/
//			case "URI":
//				states.xmedia.URI = v
//			case "TYPE":
//				states.xmedia.Type = v
//			case "GROUP-ID":
//				states.xmedia.GroupID = v
//			case "LANGUAGE":
//				states.xmedia.Language = v
//			case "NAME":
//				states.xmedia.Name = v
//			case "DEFAULT":
//				states.xmedia.Default = v
//			case "AUTOSELECT":
//				states.xmedia.Autoselect = v
//			}
//		}
//	case strings.HasPrefix(line, "#EXT-X-VERSION:"):
//		states.listtype = MASTER
//		_, err := fmt.Sscanf(line, "#EXT-X-VERSION:%d", &p.version)
//		if err != nil {
//			return errors.Wrap(err, "invalid scan version")
//		}
//	case strings.HasPrefix(line, "#EXT-X-STREAM-INF:"):
//		states.streamInf = new(VariantAttributes)
//		states.listtype = MASTER
//		for i, v := range parseLine(line[len("#EXT-X-STREAM-INF:"):]) {
//			switch i {
//			/*
//				type VariantAttributes struct {
//					Bandwidth  uint64
//					ProgramID  uint64
//					Codec      string
//					Resolution string
//					Audio      string
//					Video      string
//				}
//			*/
//			case "BANDWIDTH":
//				s, err := strconv.ParseUint(v, 10, 64)
//				if err != nil {
//					return errors.Wrap(err, "convert err of strconv BANDWIDTH")
//				}
//				states.streamInf.Bandwidth = s
//			case "PROGRAM-ID":
//				s, err := strconv.ParseUint(v, 10, 64)
//				if err != nil {
//					return errors.Wrap(err, "convert err of strconv PROGRAM-ID")
//				}
//				states.streamInf.ProgramID = s
//			case "CODECS":
//				states.streamInf.Codec = v
//			case "RESOLUTION":
//				states.streamInf.Resolution = v
//			case "AUDIO":
//				states.streamInf.Audio = v
//			case "VIDEO":
//				states.streamInf.Video = v
//			}
//		}
//	case strings.HasPrefix(line, "#EXT-X-I-FRAME-STREAM-INF:"):
//	}
//	return nil
//
//}
//
//func (p *MediaPlaylist) decodePlaylist(states *States, line string) error {
//	switch {
//	case line == "#EXTM3U":
//		states.m3u = true
//	case strings.HasPrefix(line, "#EXT-X-TARGETDURATION:"):
//		states.listtype = MEDIA
//		_, err := fmt.Sscanf(line, "#EXT-X-TARGETDURATION:%f", &p.TargetDuration)
//		if err != nil {
//			return errors.Wrap(err, " #EXT-X-TARGETDURATION scanf err")
//		}
//	case strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE:"):
//		states.listtype = MEDIA
//		_, err := fmt.Sscanf(line, "#EXT-X-MEDIA-SEQUENCE:%d", &p.MediaSequence)
//		if err != nil {
//			return errors.Wrap(err, "#EXT-X-MEDIA-SEQUENCE sanf err")
//		}
//	case strings.HasPrefix(line, "#EXT-X-VERSION:"):
//		states.listtype = MEDIA
//		_, err := fmt.Sscanf(line, "#EXT-X-VERSION:%d", &p.version)
//		if err != nil {
//			return errors.Wrap(err, "invalid scan version")
//		}
//	case strings.HasPrefix(line, "EXT-X-PLAYLIST-TYPE:"):
//		states.listtype = MEDIA
//		var playlisttype string
//		_, err := fmt.Sscanf(line, "EXT-X-PLAYLIST-TYPE:%s", &playlisttype)
//		if err != nil {
//			return errors.Wrap(err, "EXT-X-PLAYLIST-TYPE")
//		}
//		switch playlisttype {
//		case "VOD":
//			p.PlaylistType = VOD
//		case "EVENT":
//			p.PlaylistType = EVENT
//		default:
//			return errors.New("playlist type is invalid")
//		}
//	case strings.HasPrefix(line, "#EXT-X-KEY:"):
//		states.listtype = MEDIA
//		states.key = new(Key)
//		for i, v := range parseLine(line[len("#EXT-X-KEY:"):]) {
//			switch i {
//			/*
//				type Key struct {
//					Method string
//					IV     string // Initialization Vector
//					URI    string
//				}
//			*/
//			case "METHOD":
//				states.key.Method = v
//			case "IV":
//				states.key.IV = v
//			case "URI":
//				states.key.URI = v
//			}
//		}
//		states.existKey = true
//	case line == "#EXT-X-ENDLIST":
//		states.listtype = MEDIA
//	}
//	return nil
//}

func parseLine(line string) map[string]string {
	m := map[string]string{}
	lines := strings.Split(line, ",")

	for _, v := range lines {
		s := strings.Split(v, "=")
		m[s[0]] = strings.Trim(s[1], `"`)
	}

	return m
}
