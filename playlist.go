package m3u8

import (
	"fmt"
	"strings"
)

func NewPlaylist() *Playlist {
	pl := new(Playlist)
	pl.Version = uint8(3)
	pl.TargetDuration = 10
	pl.live = true

	return pl
}

func (p *Playlist) scanLineValue(line string, tag PLAYLIST) error {
	_, err := fmt.Sscanf(line, string(tag)+":%s", &p.PlaylistType)
	if err != nil {
		return err
	}
	return nil
}

func (p *Playlist) Master() bool {
	return p.master
}

func (p *Playlist) String() string {
	var b strings.Builder

	b.WriteString(string(EXTM3U) + "\n")
	if p.Master() {
		if p.hasVersion {
			b.WriteString(fmt.Sprintf("%s:%d\n", EXTVERSION, p.Version))
		}
	} else { // media playlist
		if p.hasVersion {
			b.WriteString(fmt.Sprintf("%s:%d\n", EXTVERSION, p.Version))
		}
		if p.PlaylistType != "" {
			b.WriteString(fmt.Sprintf("%s:%s\n", ExtPlaylistType, p.PlaylistType))
		}
		if p.IFrameOnly {
			b.WriteString(string(ExtIFramesOnly) + "\n")
		}
		b.WriteString(fmt.Sprintf("%s:%v\n", ExtTargetDutation, p.TargetDuration))
		b.WriteString(fmt.Sprintf("%s:%d\n", ExtMediaSequence, p.MediaSequence))
		if p.DiscontinuitySequence != 0 {
			b.WriteString(fmt.Sprintf("%s:%d\n", ExtDiscontinuitySequence, p.DiscontinuitySequence))
		}
		if p.AllowCache {
			b.WriteString(string(ExtAllowCache) + "\n")
		}
	}

	for _, v := range p.Segments {
		b.WriteString(v.String() + "\n")
	}

	if !p.live && !p.Master() {
		b.WriteString(string(EXTENDLIST))
	}

	return b.String()
}

func (p *Playlist) AppendSegment(pseg PlaylistSegment) {
	p.Segments = append(p.Segments, pseg)
}
