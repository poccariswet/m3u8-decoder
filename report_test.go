package m3u8_test

import (
	"testing"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	line := `#EXT-X-RENDITION-REPORT:URI="../1M/waitForMSN.php",LAST-MSN=273,LAST-PART=2`

	report, err := m3u8.NewReport(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "../1M/waitForMSN.php", report.URI)
	assert.Equal(t, uint64(273), report.LastMSN)
	assert.Equal(t, uint64(2), report.LastPART)
	assert.Equal(t, line, report.String())
}
