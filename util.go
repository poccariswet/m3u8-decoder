package m3u8

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//使われていないので消した方がいい
//func trimLine(line, trim string) string {
//	val := strings.TrimLeft(line, trim)
//	val = strings.Trim(val, "\n")
//	return val
//}
// mapを返すという意味合いを関数名に持たせたほうがいい、使用者はそれを意識して、プログラムを組むから
func parseLine(line string) map[string]string {
	m := map[string]string{}
	lines := strings.Split(line, ",")

	// if val has multiple items value, map's tmp key put in the value
	var tmp string
	for _, v := range lines {
		// ここのvalが決定するまでのお処理はひとまとまりとなっているため、別の関数にするべき
		// vに大して再代入が走っているので、それも防止したい
		val := validMinimumUnit(v)
		// 2がマジックナンバーになっているので、定数で定義した方が良い
		// ガード節を利用して、ネストを浅くした方がみとうしが良いのと、正常系の処理に集中できる
		// ここの条件式も式として定義しておいたほうがいい
		if len(val) != 2 {
			str := m[tmp]
			m[tmp] = fmt.Sprintf("%s,%s", str, strings.Trim(val[0], `"`))
			continue
		}

		tmp = val[0]
		m[val[0]] = strings.Trim(val[1], `"`)

	}

	return m
}

func trimLineFeed(s string) string {
	return strings.Trim(s, "\n")
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func deleteUnnecessaryString(s string) string {
	return trimSpace(trimLineFeed(s))
}

func validMinimumUnit(s string) []string {
	ns := deleteUnnecessaryString(s)
	return strings.Split(ns, "=")
}

func parseBool(line string) bool {
	// YESも定数で定義したい
	if line == "YES" {
		return true
	} else {
		return false
	}
	// 一行に直せる
	// return line == "YES"
}

// The date/time representation is ISO/IEC 8601:2004 [ISO_8601]
func parseFullTime(line string) (time.Time, error) {
	// ここでひとまとまりにしているのは、どんなフォーマットが来るかランダムだから？
	// もし、各フォーマットが区別できているのであれば、まとめるべきではない、暗黙的にパースしてしまっているから
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
	// hasでもいいけど、okって書くことのが多い気がする
	// というかいるのか？
	v, has := item[param]
	if !has {
		//not found では？
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

	// ここはparseBoolを呼び出した方が良い
	if v == "YES" {
		return true, nil
	} else {
		return false, nil
	}
}
