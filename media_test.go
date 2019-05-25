package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestMedia(t *testing.T) {
	line := `#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID="audio",LANGUAGE="eng",
ASSOC-LANGUAGE="spoken",NAME="English",AUTOSELECT=YES, FORCED=YES,
DEFAULT=YES, URI="eng/prog_index.m3u8", INSTREAM-ID="SERVICE3",
CHARACTERISTICS="public.accessibility.transcribes-spoken-dialog",
CHANNELS="10"`

	media, err := m3u8.NewMedia(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "AUDIO", media.Type)
	assert.Equal(t, "audio", media.GroupID)
	assert.Equal(t, "eng", media.Language)
	assert.Equal(t, "spoken", media.AssocLanguage)
	assert.Equal(t, "English", media.Name)
	assert.Equal(t, true, media.Autoselect)
	assert.Equal(t, true, media.Forced)
	assert.Equal(t, true, media.Default)
	assert.Equal(t, "eng/prog_index.m3u8", media.URI)
	assert.Equal(t, "SERVICE3", media.InstreamID)
	assert.Equal(t, "public.accessibility.transcribes-spoken-dialog", media.Characteristics)
	assert.Equal(t, "10", media.Channels)
}
