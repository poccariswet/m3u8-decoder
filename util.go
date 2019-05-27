package m3u8

import (
	"strconv"
	"strings"
	"time"
)

func trimLine(line, trim string) string {
	val := strings.TrimLeft(line, trim)
	val = strings.Trim(val, "\n")
	return val
}

func parseLine(line string) map[string]string {
	m := map[string]string{}
	lines := strings.Split(line, ",")

	for _, v := range lines {
		v = strings.Trim(v, "\n")
		v = strings.TrimSpace(v)
		s := strings.Split(v, "=")

		m[s[0]] = strings.Trim(s[1], `"`)
	}

	return m
}

// The date/time representation is ISO/IEC 8601:2004 [ISO_8601]
func parseFullTime(line string) (time.Time, error) {
	layouts := []string{
		"2006-01-02T15:04:05.999999999Z0700",
		time.RFC3339Nano,
		"2006-01-02T15:04:05.999999999Z07",
	}
	var (
		err error
		t   time.Time
	)
	for _, layout := range layouts {
		if t, err = time.Parse(layout, line); err == nil {
			return t, nil
		}
	}
	return t, err
}

// extract value in item, and the value parse uint64
func extractUint64(item map[string]string, param string) (uint64, error) {
	v, has := item[param]
	if !has {
		return 0, nil
	}

	uv, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}

	return uv, nil
}

// extract value in item, and the value parse float64
func extractFloat64(item map[string]string, param string) (float64, error) {
	v, has := item[param]
	if !has {
		return 0, nil
	}

	fv, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	return fv, nil
}

func extractBool(item map[string]string, param string) (bool, error) {
	v, has := item[param]
	if !has {
		return false, nil
	}

	if v == "YES" {
		return true, nil
	} else {
		return false, nil
	}
}
