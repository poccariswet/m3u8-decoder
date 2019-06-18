package m3u8

type ServerControlSegment struct{}

func NewServerControl() (*ServerControlSegment, error) {
	return &ServerControlSegment{}, nil
}

func (ss *ServerControlSegment) String() string {
	return "ServerControlSegment"
}
