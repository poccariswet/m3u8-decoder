package m3u8_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/pkg/errors"
	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestDecodeMaster(t *testing.T) {
	master, err := os.Open("./example/playlist/master.m3u8")
	if err != nil {
		t.Fatal(err)
	}

	p, err := m3u8.DecodeFrom(bufio.NewReader(master))
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, p.Master())
	assert.Equal(t, uint8(3), p.Version)
	assert.Equal(t, "", p.PlaylistType)
	assert.False(t, p.AllowCache)
	assert.Equal(t, uint64(0), p.MediaSequence)
	assert.Equal(t, float64(10), p.TargetDuration)
	assert.False(t, p.IndependentSegments)
	assert.False(t, p.IFrameOnly)
	assert.False(t, p.Discontinty)

	if len(p.Segments) == 0 {
		t.Fatal(errors.New("playlist segment is empty"))
	}

	v := p.Segments[0].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(300000), v.Bandwidth)
	assert.Equal(t, "chunklist-b300000.m3u8", v.URI)

	v = p.Segments[1].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(600000), v.Bandwidth)
	assert.Equal(t, "chunklist-b600000.m3u8", v.URI)

	v = p.Segments[2].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(850000), v.Bandwidth)
	assert.Equal(t, "chunklist-b850000.m3u8", v.URI)

	v = p.Segments[3].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(1000000), v.Bandwidth)
	assert.Equal(t, "chunklist-b1000000.m3u8", v.URI)

	v = p.Segments[4].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(1500000), v.Bandwidth)
	assert.Equal(t, "chunklist-b1500000.m3u8", v.URI)

	p, err = m3u8.ReadFile("./example/playlist/master.m3u8")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, p.Master())
	assert.Equal(t, uint8(3), p.Version)
	assert.Equal(t, "", p.PlaylistType)
	assert.False(t, p.AllowCache)
	assert.Equal(t, uint64(0), p.MediaSequence)
	assert.Equal(t, float64(10), p.TargetDuration)
	assert.False(t, p.IndependentSegments)
	assert.False(t, p.IFrameOnly)
	assert.False(t, p.Discontinty)

	if len(p.Segments) == 0 {
		t.Fatal(errors.New("playlist segment is empty"))
	}

	v = p.Segments[0].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(300000), v.Bandwidth)
	assert.Equal(t, "chunklist-b300000.m3u8", v.URI)

	v = p.Segments[1].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(600000), v.Bandwidth)
	assert.Equal(t, "chunklist-b600000.m3u8", v.URI)

	v = p.Segments[2].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(850000), v.Bandwidth)
	assert.Equal(t, "chunklist-b850000.m3u8", v.URI)

	v = p.Segments[3].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(1000000), v.Bandwidth)
	assert.Equal(t, "chunklist-b1000000.m3u8", v.URI)

	v = p.Segments[4].(*m3u8.VariantSegment)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, uint32(1500000), v.Bandwidth)
	assert.Equal(t, "chunklist-b1500000.m3u8", v.URI)
}

func TestDecodeMedia(t *testing.T) {
	media, err := os.Open("./example/playlist/media.m3u8")
	if err != nil {
		t.Fatal(err)
	}

	p, err := m3u8.DecodeFrom(bufio.NewReader(media))
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, p.Master())
	assert.Equal(t, float64(10), p.TargetDuration)

	v := p.Segments[0].(*m3u8.InfSegment)
	assert.Equal(t, float64(9.009), v.Duration)
	assert.Equal(t, "http://media.example.com/first.ts", v.URI)

	v = p.Segments[1].(*m3u8.InfSegment)
	assert.Equal(t, float64(9.009), v.Duration)
	assert.Equal(t, "http://media.example.com/second.ts", v.URI)

	v = p.Segments[2].(*m3u8.InfSegment)
	assert.Equal(t, float64(3.003), v.Duration)
	assert.Equal(t, "http://media.example.com/third.ts", v.URI)
}
