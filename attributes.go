package m3u8

import (
	"bytes"
	"time"
)

/*
	sample m3u8 playlist

	8.2.  Simple Playlist file

   #EXTM3U
   #EXT-X-TARGETDURATION:5220
   #EXTINF:5220,
   http://media.example.com/entire.ts
   #EXT-X-ENDLIST

	8.3.  Sliding Window Playlist, using HTTPS

   #EXTM3U
   #EXT-X-TARGETDURATION:8
   #EXT-X-MEDIA-SEQUENCE:2680

   #EXTINF:8,
   https://priv.example.com/fileSequence2680.ts
   #EXTINF:8,
   https://priv.example.com/fileSequence2681.ts
   #EXTINF:8,
   https://priv.example.com/fileSequence2682.ts

	8.4.  Playlist file with encrypted media segments

   #EXTM3U
   #EXT-X-MEDIA-SEQUENCE:7794
   #EXT-X-TARGETDURATION:15

   #EXT-X-KEY:METHOD=AES-128,URI="https://priv.example.com/key.php?r=52"

   #EXTINF:15,
   http://media.example.com/fileSequence52-1.ts
   #EXTINF:15,
   http://media.example.com/fileSequence52-2.ts
   #EXTINF:15,
   http://media.example.com/fileSequence52-3.ts

   #EXT-X-KEY:METHOD=AES-128,URI="https://priv.example.com/key.php?r=53"

   #EXTINF:15,
   http://media.example.com/fileSequence53-1.ts

	8.5.  Variant Playlist file

   #EXTM3U
   #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1280000
   http://example.com/low.m3u8
   #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=2560000
   http://example.com/mid.m3u8
   #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=7680000
   http://example.com/hi.m3u8
   #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=65000,CODECS="mp4a.40.5"
   http://example.com/audio-only.m3u8

	8.6.  Variant Playlist with I-Frames

   In this example, the PROGRAM-ID attributes have been left out:

   #EXTM3U
   #EXT-X-STREAM-INF:BANDWIDTH=1280000
   low/audio-video.m3u8
   #EXT-X-I-FRAME-STREAM-INF:BANDWIDTH=86000,URI="low/iframe.m3u8"
   #EXT-X-STREAM-INF:BANDWIDTH=2560000
   mid/audio-video.m3u8
   #EXT-X-I-FRAME-STREAM-INF:BANDWIDTH=150000,URI="mid/iframe.m3u8"
   #EXT-X-STREAM-INF:BANDWIDTH=7680000
   hi/audio-video.m3u8
   #EXT-X-I-FRAME-STREAM-INF:BANDWIDTH=550000,URI="hi/iframe.m3u8"
   #EXT-X-STREAM-INF:BANDWIDTH=65000,CODECS="mp4a.40.5"
   audio-only.m3u8

	8.7.  Variant Playlist with Alternative audio

   #EXTM3U
   #EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID="aac",NAME="English", \
      DEFAULT=YES,AUTOSELECT=YES,LANGUAGE="en", \
      URI="main/english-audio.m3u8"
   #EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID="aac",NAME="Deutsche", \
      DEFAULT=NO,AUTOSELECT=YES,LANGUAGE="de", \
      URI="main/german-audio.m3u8"
   #EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID="aac",NAME="Commentary", \
      DEFAULT=NO,AUTOSELECT=NO,URI="commentary/audio-only.m3u8"
   #EXT-X-STREAM-INF:BANDWIDTH=1280000,CODECS="...",AUDIO="aac"
   low/video-only.m3u8
   #EXT-X-STREAM-INF:BANDWIDTH=2560000,CODECS="...",AUDIO="aac"
   mid/video-only.m3u8
   #EXT-X-STREAM-INF:BANDWIDTH=7680000,CODECS="...",AUDIO="aac"
   hi/video-only.m3u8
   #EXT-X-STREAM-INF:BANDWIDTH=65000,CODECS="mp4a.40.5",AUDIO="aac"
   main/english-audio.m3u8

*/

type PlaylistSegment interface {
	String() string
}

type ListType int

const (
	ERRTYPE ListType = iota
	MASTER
	MEDIA
)

// EXT-X-PLAYLIST-TYPE tag with a value of either EVENT or VOD
type PlaylistType uint

const (
	EVENT PlaylistType = iota
	VOD
)

// For decrypt media segments
type Key struct {
	Method string
	IV     string // Initialization Vector
	URI    string
}

// EXT-X-MAP
type Map struct {
	URI        string
	ByteRangeN int64 // EXT-X-BYTERANGE uri length
	ByteRangeO int64 // EXT-X-BYTERANGE uri offset
}

// The EXT-X-MEDIA tag is used to relate Playlists that contain alternative renditions of the same content.
type XMedia struct {
	URI        string
	Type       string
	GroupID    string
	Language   string
	Name       string
	Default    string
	Autoselect string
}

// media segment in media playlist
// If the Playlist file does not contain an EXT-X-KEY tag then media segments are not encrypted.
type MediaSeqment struct {
	SequenceNum     uint64
	Title           string
	ByteRangeN      int64 // EXT-X-BYTERANGE uri length
	ByteRangeO      int64 // EXT-X-BYTERANGE uri offset
	URI             string
	Duration        float64   // EXTINF
	Key             *Key      // EXT-X-KEY
	ProgramDateTime time.Time // EXT-X-PROGRAM-DATE-TIME:<YYYY-MM-DDThh:mm:ssZ>
	Map             *Map      //EXT-X-MAP
	Discontinuity   bool      // EXT-X-DISCONTINUITY encoding discontinuity between the media segment that follows it and the one that preceded it.
}

type MediaPlaylist struct {
	buf          bytes.Buffer
	capacity     uint // max capacity of playlist
	playlistSize uint // playlist size of mediasegment

	MediaSeqments  []*MediaSeqment
	MediaSequence  uint64
	version        uint8
	TargetDuration float64
	AllowCache     bool         // EXT-X-ALLOW-CACHE:<YES|NO>
	PlaylistType   PlaylistType // EXT-X-PLAYLIST-TYPE:<EVENT|VOD>
}

type VariantAttributes struct {
	Bandwidth  uint64
	ProgramID  uint64
	Codec      string
	Resolution string
	Audio      string
	Video      string
}

type VariantPlaylist struct {
	URI            string
	MediaPlaylists *MediaPlaylist
	VariantAttributes
}

type MasterPlaylist struct {
	buf bytes.Buffer

	version          uint8
	VariantPlaylists []*VariantPlaylist
}

type Playlist struct {
	version        uint8
	PlaylistType   PlaylistType
	AllowCache     bool
	MediaSequence  uint64
	TargetDuration float64
	IFrameOnly     bool // EXT-X-I-FRAMES-ONLY
	master         bool

	Segments []PlaylistSegment
}

// state of m3u and temporary store segments, stream inf...etc
type States struct {
	m3u       bool
	listtype  ListType
	xmedia    XMedia
	streamInf *VariantAttributes
	existKey  bool
	key       *Key
}
