package m3u8

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

func (p *MasterPlaylist) Init() error {
	p.version = uint8(3) // most min version
	return nil
}

func (p *MediaPlaylist) Init() error {
	p.version = uint8(3)
	p.capacity = 1024
	p.playlistSize = 8
	p.MediaSeqments = make([]*MediaSeqment, p.capacity)

	return nil
}

func (p *MasterPlaylist) DecodeFrom(r io.Reader) error {
	return nil
}

func (p *MediaPlaylist) DecodeFrom(r io.Reader) error {
	return nil
}

func (p *MasterPlaylist) String() string {
	return ""
}

func (p *MediaPlaylist) String() string {
	return ""
}

func decode(buf *bytes.Buffer) (Playlist, ListType, error) {
	var master *MasterPlaylist
	var media *MediaPlaylist
	var listtype ListType
	var end bool
	var states *States

	if err := master.Init(); err != nil {
		return nil, ERRTYPE, err
	}

	if err := media.Init(); err != nil {
		return nil, ERRTYPE, err
	}

	for !end {
		line, err := buf.ReadString('\n')
		if err != io.EOF {
			end = true
		} else if err != nil {
			break
		}

		if len(line) < 1 || line == "\r" {
			continue
		}

		line = strings.TrimSpace(line)
		if err := decodeMasterPlaylist(master, states, listtype, line); err != nil {
			return master, listtype, err
		}

		if err := decodeMediaPlaylist(media, states, listtype, line); err != nil {
			return media, listtype, err
		}
	}

	switch listtype {
	case MASTER:
		return master, listtype, nil
	case MEDIA:
		return media, listtype, nil
	}

	return nil, 2, errors.New("not playlist")
}

func DecodeFrom(r io.Reader) (Playlist, ListType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, 2, err
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
	}
	return nil

}

func decodeMediaPlaylist(playlist *MediaPlaylist, states *States, listtype ListType, line string) error {
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
