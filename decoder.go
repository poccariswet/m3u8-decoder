package m3u8

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

func NewMasterPlaylist() *MasterPlaylist {
	master := new(MasterPlaylist)
	master.version = uint8(3) // most min version

	return master
}

func NewMediaPlaylist() *MediaPlaylist {
	media := new(MediaPlaylist)
	media.version = uint8(3) // most min version
	media.capacity = 1024
	media.playlistSize = 8
	media.MediaSeqments = make([]*MediaSeqment, media.capacity)

	return media
}

func (p *MasterPlaylist) DecodeFrom(r io.Reader) error {
	return nil
}

func (p *MediaPlaylist) DecodeFrom(r io.Reader) error {
	return nil
}

func (p *MasterPlaylist) String() string {
	return "MASTER PLAYLIST"
}

func (p *MediaPlaylist) String() string {
	return "MEDIA PLAYLIST"
}

func decode(buf *bytes.Buffer) (Playlist, ListType, error) {
	master := NewMasterPlaylist()
	media := NewMediaPlaylist()
	var listtype ListType
	var end bool
	states := new(States)

	for !end {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			end = true
		} else if err != nil {
			break
		}

		if len(line) < 1 || line == "\r" {
			continue
		}

		line = strings.TrimSpace(line)
		if err := decodeMasterPlaylist(master, states, listtype, line); err != nil {
			return master, ERRTYPE, err
		}

		if err := decodeMediaPlaylist(media, states, listtype, line); err != nil {
			return media, ERRTYPE, err
		}
	}

	switch listtype {
	case MASTER:
		return master, listtype, nil
	case MEDIA:
		return media, listtype, nil
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

func decodeMasterPlaylist(playlist *MasterPlaylist, states *States, listtype ListType, line string) error {
	switch {
	case line == "#EXTM3U":
		states.m3u = true
	case strings.HasPrefix(line, "#EXT-X-MEDIA:"):
		listtype = MASTER
		for i, v := range parseLine(line[len("#EXT-X-MEDIA:"):]) {
			switch i {
			/*
				type XMedia struct {
					URI        string
					Type       string
					GroupID    string
					Language   string
					Name       string
					Default    string
					Autoselect string
				}
			*/
			case "URI":
				states.xmedia.URI = v
			case "TYPE":
				states.xmedia.Type = v
			case "GROUP-ID":
				states.xmedia.GroupID = v
			case "LANGUAGE":
				states.xmedia.Language = v
			case "NAME":
				states.xmedia.Name = v
			case "DEFAULT":
				states.xmedia.Default = v
			case "AUTOSELECT":
				states.xmedia.Autoselect = v
			}
		}
	case strings.HasPrefix(line, "#EXT-X-VERSION:"):
		listtype = MASTER
		_, err := fmt.Sscanf(line, "#EXT-X-VERSION:%s", &playlist.version)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, "#EXT-X-STREAM-INF:"):
		listtype = MASTER
		for i, v := range parseLine(line[len("#EXT-X-STREAM-INF:"):]) {
			switch i {
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
			case "BANDWIDTH":
				states.streamInf.Bandwidth = v
			case "PROGRAM-ID":
				states.streamInf.ProgramID = v
			case "CODECS":
				states.streamInf.Codec = v
			case "RESOLUTION":
				states.streamInf.Resolution = v
			case "AUDIO":
				states.streamInf.Audio = v
			case "VIDEO":
				states.streamInf.Video = v
			}
		}
		return nil
	}
	return nil

}

func decodeMediaPlaylist(playlist *MediaPlaylist, states *States, listtype ListType, line string) error {
	switch line {
	case line == "#EXTM3U":
		states.m3u = true
	}
	return nil
}

func parseLine(line string) map[string]string {
	m := map[string]string{}
	lines := strings.Split(line, ",")

	for _, v := range lines {
		s := strings.Split(v, "=")
		m[s[0]] = strings.Trim(s[1], `"`)
	}

	return m
}
