package dcdkatsu

import (
	"regexp"
	"time"
)

func parseDataGetDate(dataGetDate string) (time.Time, error) {
	re := regexp.MustCompile("データ取得日：(.*)$")
	dgd := re.ReplaceAllString(
		dataGetDate,
		"$1",
	)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.ParseInLocation("2006年01月02日 15時04分", dgd, loc)
}
