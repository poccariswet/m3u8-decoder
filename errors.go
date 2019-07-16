package m3u8

import "github.com/pkg/errors"

var (
	InvalidPlaylist = errors.New("invalid playlist, must start with #EXTM3U")

	InvalidPlaylistType = errors.New("invalid playlist, mixed master and media")

	NotFound = errors.New("not found err")
)
