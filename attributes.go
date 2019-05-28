package m3u8

import "time"

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

// EXT-X-PROGRAM-DATE-TIME tag segment
type DateTimeSegment struct {
	Time time.Time
}

// For decrypt media segments
type KeySegment struct {
	Method            string
	IV                string // Initialization Vector
	URI               string
	KeyFormat         string
	KeyFormatVersions string
}

// EXT-X-MAP segment
type MapSegment struct {
	URI       string
	ByteRange *ByteRangeSegment
}

// The EXT-X-MEDIA tag is used to relate Playlists that contain alternative renditions of the same content.
type MediaSegment struct {
	Type            string
	GroupID         string
	Language        string
	AssocLanguage   string
	Name            string
	Autoselect      bool
	Forced          bool
	Default         bool
	URI             string
	InstreamID      string
	Characteristics string
	Channels        string
}

// EXT-X-DATERANGE range of time defined by a starting and ending date with a set of attribute / value pairs
type DateRangeSegment struct {
	ID              string
	Class           string
	StartDate       time.Time
	EndDate         time.Time
	Duration        float64
	PlannedDuration float64
	Scte35Cmd       string
	Scte35Out       string
	Scte35In        string
	EndOnNext       bool
}

type ByteRangeSegment struct {
	Length int64 // the length of the sub-range in bytes
	Offset int64 // a byte offset from the beginning of the resource
}

// #EXTINF attribute and, under uri
type InfSegment struct {
	Duration float64
	URI      string
}

// #EXT-X-STREAM-INF or EXT-X-I-FRAME-STREAM-INF attributes
type VariantSegment struct {
	IFrame bool

	URI              string
	Bandwidth        uint32
	Name             string
	Subtitle         string
	AverageBandwidth uint32
	ProgramID        uint32
	Codec            string
	AudioCodec       string
	Audio            string
	Video            string
	FrameRate        float64
	ClosedCaptions   string
	HDCPLevel        string
	Resolution       *Resolution
}

type Playlist struct {
	version               uint8
	PlaylistType          string
	AllowCache            bool
	MediaSequence         uint64
	DiscontinuitySequence uint64
	TargetDuration        float64
	IFrameOnly            bool // EXT-X-I-FRAMES-ONLY
	master                bool
	live                  bool
	Discontinty           bool // EXT-X-DISCONTINUITY encoding discontinuity between the media segment that follows it and the one that preceded it.

	Segments []PlaylistSegment
}

// state of m3u and temporary store segments, stream inf...etc
type States struct {
	m3u8       bool
	master     bool
	segmentTag bool
	listtype   ListType
	segment    PlaylistSegment
}
