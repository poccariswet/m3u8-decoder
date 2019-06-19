package m3u8_test

import (
	"strings"
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestServerContorol(t *testing.T) {
	line := `#EXT-X-SERVER-CONTROL:CAN-BLOCK-RELOAD=YES,
PART-HOLD-BACK=1.5,CAN-SKIP-UNTIL=12.5,HOLD-BACK=6.5`

	sc, err := m3u8.NewServerControl(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, sc.CanBlockReload)
	assert.Equal(t, float64(12.5), sc.CanSkipUntil)
	assert.Equal(t, float64(6.5), sc.HoldBack)
	assert.Equal(t, float64(1.5), sc.PartHoldBack)
	assert.Equal(t, strings.Replace(line, "\n", "", 1), sc.String())
}
