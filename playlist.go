package m3u8

func NewPlaylist() *Playlist {
	pl := new(Playlist)
	pl.version = uint8(3)
	pl.TargetDuration = 10
	pl.live = true

	return pl
}

func (p *Playlist) Master() bool {
	return p.master
}

// TODO display playlist contents
func (p *Playlist) String() string {
	if p.master {
		return "MASTER PLAYLIST"
	}
	return "MEDIA PLAYLIST"
}

func (p *Playlist) AppendSegment(pseg PlaylistSegment) {
	p.Segments = append(p.Segments, pseg)
}
