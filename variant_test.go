package m3u8_test
// ここだけパッケージ名を分けているのは？
import (
	"strings"
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

// testifyの内部でDeepEqualが使用されているので、あまり使わない方がいい
// github.com/google/go-cmp/cmpとかを使う方が安全
func TestVariant(t *testing.T) {
	line := `#EXT-X-I-FRAME-STREAM-INF:AVERAGE-BANDWIDTH=305152,BANDWIDTH=435408,
AUDIO="sample",VIDEO="sample",CODECS="mp4a.40.2",
RESOLUTION=1280x720,FRAME-RATE=24.001,CLOSED-CAPTIONS=NONE,
PROGRAM-ID=1,NAME="1280p",HDCP-LEVEL=TYPE-0,SUBTITLES="sample.subs",URI="sample.url"`

	l := line[len(m3u8.ExtFrameStreamInf+":"):]
	v, err := m3u8.NewVariant(l)
	if err != nil {
		t.Fatal(err)
	}
	v.IFrame = true

	assert.Nil(t, err)
	assert.Equal(t, "sample.url", v.URI)
	assert.Equal(t, uint32(435408), v.Bandwidth)
	assert.Equal(t, "1280p", v.Name)
	assert.Equal(t, "sample.subs", v.Subtitle)
	assert.Equal(t, uint32(305152), v.AverageBandwidth)
	assert.Equal(t, uint32(1), v.ProgramID)
	assert.Equal(t, "mp4a.40.2", v.Codecs)
	assert.Equal(t, "sample", v.Audio)
	assert.Equal(t, "sample", v.Video)
	assert.Equal(t, float64(24.001), v.FrameRate)
	assert.Equal(t, "NONE", v.ClosedCaptions)
	assert.Equal(t, "TYPE-0", v.HDCPLevel)

	assert.NotNil(t, v.Resolution)
	assert.Equal(t, uint16(1280), v.Resolution.Width)
	assert.Equal(t, uint16(720), v.Resolution.Height)
	assert.Equal(t, strings.Replace(line, "\n", "", 4), v.String())
}
