package m3u8

import (
	"bytes"
	"io"
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

func decode(buf bytes.Buffer) (Playlist, ListType, error) {
	var master *MasterPlaylist
	var media *MediaPlaylist
	var listtype ListType

	if err := master.Init(); err != nil {
		return nil, ERRTYPE, err
	}

	if err := media.Init(); err != nil {
		return nil, ERRTYPE, err
	}

	// ↓ read していく

	return nil, nil, nil
}

func DecodeFrom(r io.Reader) (Playlist, ListType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, 2, err
	}
	return decode(buf)
}
