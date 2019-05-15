package m3u8

const (

	// Playlist
	EXTM3U                   = `#EXTM3U`
	ExtENDLIST               = `#EXT-X-ENDLIST`
	ExtTargetDutation        = `#EXT-X-TARGETDURATION`
	ExtAllowCache            = `#EXT-X-ALLOW-CACHE`
	ExtDiscontinuitySequence = `#EXT-X-DISCONTINUITY-SEQUENCE`
	ExtIndependentSegments   = `#EXT-X-INDEPENDENT-SEGMENTS`
	ExtPlaylistType          = `#EXT-X-PLAYLIST-TYPE`
	ExtIFramesOnly           = `#EXT-X-I-FRAMES-ONLY`
	ExtMediaSequence         = `#EXT-X-MEDIA-SEQUENCE`
	ExtVersion               = `#EXT-X-VERSION`

	// ByteRange
	ByteRange = "BYTERANGE"

	// Encrypt
	METHOD            = "METHOD"
	URI               = "URI"
	IV                = "IV"
	KEYFORMAT         = "KEYFORMAT"
	KeyFormatVersions = "KEYFORMATVERSIONS"

	// DateRange
	ID              = "ID"
	CLASS           = "CLASS"
	STARTDATE       = "START-DATE"
	ENDDATE         = "END-DATE"
	DURATION        = "DURATION"
	PLANNEDDURATION = "PLANNED-DURATION"
	SCTE35CMD       = "SCTE35-CMD"
	SCTE35OUT       = "SCTE35-OUT"
	SCTE35IN        = "SCTE35-IN"
	ENDONNEXT       = "END-ON-NEXT"

	// Ext-Start
	TIMEOFFSET = "TIME-OFFSET"
	PRECISE    = "PRECISE"

	// Session
	DATAID   = "DATA-ID"
	VALUE    = "VALUE"
	LANGUAGE = "LANGUAGE"

	// Variant Playlist
	ExtSessionKey      = `#EXT-X-SESSION-KEY`
	ExtKey             = `#EXT-X-KEY`
	ExtDiscontinuity   = `#EXT-X-DISCONTINUITY`
	ExtProgramDateTime = `#EXT-X-PROGRAM-DATE-TIME`
	ExtDateRange       = `#EXT-X-DATERANGE`
	ExtMap             = `#EXT-X-MAP`
	ExtSessionData     = `#EXT-X-SESSION-DATA`
	ExtEXTINF          = `#EXTINF`
	ExtByteRange       = `#EXT-X-BYTERANGE`
	ExtStart           = `#EXT-X-START`
	ExtMedia           = `#EXT-X-MEDIA`
	ExtStreamInf       = `#EXT-X-STREAM-INF`
	ExtFrameStreamInf  = `#EXT-X-I-FRAME-STREAM-INF`

	// MediaType
	TYPE            = "TYPE"
	GROUPID         = "GROUP-ID"
	ASSOCLANGUAGE   = "ASSOC-LANGUAGE"
	NAME            = "NAME"
	AUTOSELECT      = "AUTOSELECT"
	DEFAULT         = "DEFAULT"
	FORCED          = "FORCED"
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
