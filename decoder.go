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

/*
m3u8
| Tag                            | Occured in | Proto ver | In Go lib since |
|--------------------------------|------------|-----------|-----------------|
| EXT-X-ALLOW-CACHE              | MED        | 1         | 0.1             |
| EXT-X-BYTERANGE                | MED        | 4         | 0.1             |
| EXT-X-DISCONTINUITY            | MED        | 1         | 0.2             |
| EXT-X-DISCONTINUITY-SEQUENCE   | MED        | 6         |                 |
| EXT-X-ENDLIST                  | MED        | 1         | 0.1             |
| EXT-X-I-FRAME-STREAM-INF       | MAS        | 4         | 0.3             |
| EXT-X-I-FRAMES-ONLY            | MED        | 4         | 0.3             |
| EXT-X-INDEPENDENT-SEGMENTS     | MAS        | 6         |                 |
| EXT-X-KEY                      | MED        | 1         | 0.1             |
| EXT-X-MAP                      | MED        | 5         | 0.3             |
| EXT-X-MEDIA                    | MAS        | 4         | 0.1             |
| EXT-X-MEDIA-SEQUENCE           | MED        | 1         | 0.1             |
| EXT-X-PLAYLIST-TYPE            | MED        | 3         | 0.2             |
| EXT-X-PROGRAM-DATE-TIME        | MED        | 1         | 0.2             |
| EXT-X-SESSION-DATA             | MAS        | 7         |                 |
| EXT-X-START                    | MAS        | 6         |                 |
| EXT-X-STREAM-INF               | MAS        | 1         | 0.1             |
| EXT-X-TARGETDURATION           | MED        | 1         | 0.1             |
| EXT-X-VERSION                  | MAS        | 2         | 0.1             |
| EXTINF                         | MED        | 1         | 0.1             |
| EXTM3U                         | MAS,MED    | 1         | 0.1             |

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
