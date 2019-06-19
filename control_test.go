package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestServerContorol(t *testing.T) {
	line := `#EXT-X-SERVER-CONTROL:CAN-BLOCK-RELOAD=YES,
PART-HOLD-BACK=1.0,CAN-SKIP-UNTIL=12.0,HOLD-BACK=6.0`

	sc, err := m3u8.NewServerControl(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, sc.CanBlockReload)
	assert.Equal(t, float64(12.0), sc.CanSkipUntil)
	assert.Equal(t, float64(6.0), sc.HoldBack)
	assert.Equal(t, float64(1.0), sc.PartHoldBack)
}
