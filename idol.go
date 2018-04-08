package dcdkatsu

import (
	"fmt"
)

// Idol hold information for Aikatsu! idol.
type Idol struct {
	Id  string
	URL string
	// Levels for each type
	CuteLevel int64
	CoolLevel int64
	SexyLevel int64
	PopLevel  int64

	FanCount     int64
	LastLocation string

	// TODO: Check DailyHighScore type; is it int?
	// If there's no rank, set to zero.
	DailyHighScoreSolo    int64
	DailyHighScoreFriends int64
}

// FetchIdol returns parsed Idol data.
// This data will be fetched from http://mypage.aikatsu.com/mypages/index/<ID> .
func FetchIdol(id string) (Idol, error) {
	idol := Idol{}
	idol.Id = id
	idol.URL = fmt.Sprintf("http://mypage.aikatsu.com/mypages/index/%s", id)
	return idol, nil
}
