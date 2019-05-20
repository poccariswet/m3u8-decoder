package m3u8

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// RESOLUTION=640x360
type Resolution struct {
	Width  uint16
	Height uint16
}

func NewResolution(val string, has bool) (*Resolution, error) {
	if !has {
		return nil, nil
	}
	s := strings.Split(val, "x")
	if len(s) == 2 {
		return nil, errors.New("invalid RESOLUTION value")
	}

	// width
	w, err := strconv.ParseUint(s[0], 10, 16)
	if err != nil {
		return nil, errors.Wrap(err, "parse uint err: ")
	}
	// height
	h, err := strconv.ParseUint(s[1], 10, 16)
	if err != nil {
		return nil, errors.Wrap(err, "parse uint err: ")
	}

	return &Resolution{
		Width:  uint16(w),
		Height: uint16(h),
	}, nil
}
