package m3u8_test

import (
	"testing"
	"time"

	m3u8 "github.com/poccariswet/m3u8-decoder"
	"github.com/stretchr/testify/assert"
)

func TestDateRange(t *testing.T) {
	line := `#EXT-X-DATERANGE:ID=splice-6FFFFFF0,CLASS=sample.class,START-DATE=2014-03-05T11:15:00Z,END-DATE=2014-03-06T11:15:00Z,DURATION=59.993,PLANNED-DURATION=59.993,SCTE35-CMD=0xFC002F0000000000FF,SCTE35-OUT=0xFC002F0000000000FF000014056FFFFFF00,SCTE35-IN=0xFC002F0000000000FF,END-ON-NEXT=YES`

	dr, err := m3u8.NewDateRange(line)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "splice-6FFFFFF0", dr.ID)
	assert.Equal(t, "sample.class", dr.Class)
	assert.Equal(t, float64(59.993), dr.Duration)
	assert.Equal(t, float64(59.993), dr.PlannedDuration)
	assert.Equal(t, "0xFC002F0000000000FF", dr.Scte35Cmd)
	assert.Equal(t, "0xFC002F0000000000FF000014056FFFFFF00", dr.Scte35Out)
	assert.Equal(t, "0xFC002F0000000000FF", dr.Scte35In)
	assert.True(t, dr.EndOnNext)

	expectedStartDate, err := time.Parse("2006-01-02T15:04:05.999999999Z07", "2014-03-05T11:15:00Z")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedStartDate, dr.StartDate)

	expectedEndDate, err := time.Parse("2006-01-02T15:04:05.999999999Z07", "2014-03-06T11:15:00Z")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedEndDate, dr.EndDate)

	assert.Equal(t, line, dr.String())
}
