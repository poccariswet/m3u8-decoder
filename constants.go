package m3u8

import (
	"strings"
)

type PLAYLIST string

func (p PLAYLIST) len() int {
	return len(p + ":")
}

func (p PLAYLIST) match(line string) bool {
	return strings.HasPrefix(line, string(p))
}

const (

	// Playlist
	EXTM3U     PLAYLIST = `#EXTM3U`
	EXTENDLIST PLAYLIST = `#EXT-X-ENDLIST`
	EXTVERSION PLAYLIST = `#EXT-X-VERSION`

	ExtTargetDutation        PLAYLIST = `#EXT-X-TARGETDURATION`
	ExtAllowCache            PLAYLIST = `#EXT-X-ALLOW-CACHE`
	ExtDiscontinuitySequence PLAYLIST = `#EXT-X-DISCONTINUITY-SEQUENCE`
	ExtIndependentSegments   PLAYLIST = `#EXT-X-INDEPENDENT-SEGMENTS`
	ExtPlaylistType          PLAYLIST = `#EXT-X-PLAYLIST-TYPE`
	ExtIFramesOnly           PLAYLIST = `#EXT-X-I-FRAMES-ONLY`
	ExtMediaSequence         PLAYLIST = `#EXT-X-MEDIA-SEQUENCE`

	// Variant Playlist
	ExtSessionKey      PLAYLIST = `#EXT-X-SESSION-KEY`
	ExtKey             PLAYLIST = `#EXT-X-KEY`
	ExtDiscontinuity   PLAYLIST = `#EXT-X-DISCONTINUITY`
	ExtProgramDateTime PLAYLIST = `#EXT-X-PROGRAM-DATE-TIME`
	ExtDateRange       PLAYLIST = `#EXT-X-DATERANGE`
	ExtMap             PLAYLIST = `#EXT-X-MAP`
	ExtSessionData     PLAYLIST = `#EXT-X-SESSION-DATA`
	EXTINF             PLAYLIST = `#EXTINF`
	ExtByteRange       PLAYLIST = `#EXT-X-BYTERANGE`
	ExtStart           PLAYLIST = `#EXT-X-START`
	ExtMedia           PLAYLIST = `#EXT-X-MEDIA`
	ExtStreamInf       PLAYLIST = `#EXT-X-STREAM-INF`
	ExtFrameStreamInf  PLAYLIST = `#EXT-X-I-FRAME-STREAM-INF`

	// for low-latency
	ExtServerControl   PLAYLIST = `#EXT-X-SERVER-CONTROL`
	ExtPartInf         PLAYLIST = `#EXT-X-PART-INF`
	ExtPart            PLAYLIST = `#EXT-X-PART`
	ExtRenditionReport PLAYLIST = `#EXT-X-RENDITION-REPORT`
	ExtSkip            PLAYLIST = `#EXT-X-SKIP`

	// rendition report
	LASTMSN  = "LAST-MSN"
	LASTPART = "LAST-PART"

	// server control
	CANBLOCKRELOAD = "CAN-BLOCK-RELOAD"
	HOLDBACK       = "HOLD-BACK"
	PARTHOLDBACK   = "PART-HOLD-BACK"
	CANSKIPUNTIL   = "CAN-SKIP-UNTIL"

	// part
	PARTTARGET  = "PART-TARGET"
	INDEPENDENT = "INDEPENDENT"
	GAP         = "GAP"

	// skip
	SKIPPEDSEGMENTS = "SKIPPED-SEGMENTS"

	// Utility tag
	URI      = "URI"
	DURATION = "DURATION"

	// ByteRange
	BYTERANGE = "BYTERANGE"

	// Encrypt key
	METHOD            = "METHOD"
	IV                = "IV"
	KEYFORMAT         = "KEYFORMAT"
	KEYFORMATVERSIONS = "KEYFORMATVERSIONS"

	// DateRange
	ID              = "ID"
	CLASS           = "CLASS"
	STARTDATE       = "START-DATE"
	ENDDATE         = "END-DATE"
	PLANNEDDURATION = "PLANNED-DURATION"
	SCTE35CMD       = "SCTE35-CMD"
	SCTE35OUT       = "SCTE35-OUT"
	SCTE35IN        = "SCTE35-IN"
	ENDONNEXT       = "END-ON-NEXT"

	// Ext-Start
	TIMEOFFSET = "TIME-OFFSET"
	PRECISE    = "PRECISE"

	// Session
	DATAID = "DATA-ID"
	VALUE  = "VALUE"

	// MediaType
	TYPE            = "TYPE"
	GROUPID         = "GROUP-ID"
	LANGUAGE        = "LANGUAGE"
	ASSOCLANGUAGE   = "ASSOC-LANGUAGE"
	NAME            = "NAME"
	AUTOSELECT      = "AUTOSELECT"
	FORCED          = "FORCED"
	DEFAULT         = "DEFAULT"
	INSTREAMID      = "INSTREAM-ID"
	CHARACTERISTICS = "CHARACTERISTICS"
	CHANNELS        = "CHANNELS"

	/// Variant
	RESOLUTION       = "RESOLUTION"
	PROGRAMID        = "PROGRAM-ID"
	CODECS           = "CODECS"
	BANDWIDTH        = "BANDWIDTH"
	AVERAGEBANDWIDTH = "AVERAGE-BANDWIDTH"
	FRAMERATE        = "FRAME-RATE"
	VIDEO            = "VIDEO"
	AUDIO            = "AUDIO"
	SUBTITLES        = "SUBTITLES"
	CLOSEDCAPTIONS   = "CLOSED-CAPTIONS"
	HDCPLEVEL        = "HDCP-LEVEL"
)
