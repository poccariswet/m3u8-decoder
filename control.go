package m3u8

// ServerControl indicate support for features such as Blocking Playlist Reload and Playlist Delta Updates
type ServerControlSegment struct {
	CanBlockReload bool
	CanSkipUntil   float64
	HoldBack       float64
	PartHoldBack   float64
}

// New server control segment
func NewServerControl() (*ServerControlSegment, error) {
	return &ServerControlSegment{}, nil
}

// segment to string
func (ss *ServerControlSegment) String() string {
	return "ServerControlSegment"
}
