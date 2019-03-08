package m3u8

import (
	"bytes"
	"io"
)

type Playlist interface {
	Encode() *bytes.Buffer
	Decode(bytes.Buffer, bool) error
	DecodeFrom(io.Reader) error
	String() string
}

type PlaylistType uint

/*
 Media Playlist file sample:

   #EXTM3U
   #EXT-X-VERSION:3
   #EXT-X-TARGETDURATION:5
   #EXTINF:5,
   http://media.jp/sound/b/20190117/20190117_010000_mgNac.aac
   #EXT-X-ENDLIST

 Sample Media Playlist, using HTTPS:
   EXTM3U
   #EXT-X-VERSION:3
   #EXT-X-ALLOW-CACHE:NO
   #EXT-X-TARGETDURATION:5
   #EXT-X-MEDIA-SEQUENCE:1
   #EXT-X-PROGRAM-DATE-TIME:2019-01-17T01:00:00+09:00
   #EXTINF:5,
   https://media.jp/sound/b/20190117/20190117_010000_mgNac.aac
   #EXT-X-PROGRAM-DATE-TIME:2019-01-17T01:00:05+09:00
   #EXTINF:5,
   https://media.jp/sound/b/20190117/20190117_010005_T8DCV.aac
   ...
*/

// https://github.com/poccariswet/m3u8/blob/master/m3u8.md
type MediaPlaylist struct {
	version        uint8
	TargetDuration float64
}

// Decode From returns playlist and find the type.
func DecodeFrom(r io.Reader) (Playlist, PlaylistType, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, 0, err
	}

	return decode(buf)
}

// decode is master media play list decoder
func decode(buf *bytes.Buffer) (PlayList, PlayListType, error) {
	var medialist MediaPlaylist
	var playlistType PlaylistType

	return nil, 0, nil
}
