package m3u8

import (
	"strconv"
	"strings"
)

func trimLine(line, trim string) string {
	val := strings.TrimLeft(line, trim)
	val = strings.Trim(v, "\n")
	return val
}

func parseLine(line string) map[string]string {
	m := map[string]string{}
	lines := strings.Split(line, ",")

	for _, v := range lines {
		s := strings.Split(v, "=")
		m[s[0]] = strings.Trim(s[1], `"`)
	}

	return m
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
