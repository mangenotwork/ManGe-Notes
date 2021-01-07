package util

import (
	"fmt"
	"strings"
	"time"
)

//FormatTime 格式化时间显示
func FormatTime(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

//DateTime DateTime
func DateTime(unix int64) (ret string) {
	if unix < 1 {
		return
	}
	tm := time.Unix(unix, 0)
	ret = FormatTime(tm, "YYYY-MM-DD HH:mm:ss")
	return
}

//DateTimeFormat DateTimeFormat
func DateTimeFormat(unix int64, format string) (ret string) {
	if unix < 1 {
		return
	}
	tm := time.Unix(unix, 0)
	ret = FormatTime(tm, format)
	return
}
