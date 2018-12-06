package dcdkatsu

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Idol hold information for Aikatsu! idol.
type Idol struct {
	ID  string
	URL string

	// Basic Info
	Name             string
	AvatarURL        string
	LastPlayLocation string

	// Levels for each type
	CuteLevel int
	CoolLevel int
	SexyLevel int
	PopLevel  int

	// Misc
	IdolRank         int
	IdolRankLabel    string
	FanCount         int
	PlayedCardsCount int

	// TODO: Check DailyHighScore type; is it int?
	// If there's no rank, set to zero.
	RankingHighscoreSolo     int
	RankingHighscoresFriends int
	RankingTotalfan          int

	// DataGetDate is not , but upstream api called time
	DataGetDate time.Time
}

// FetchIdol returns parsed Idol data.
// This data will be fetched from http://mypage.aikatsu.com/mypages/index/<ID> .
func FetchIdol(id string) (Idol, error) {
	idol := Idol{
		ID:  id,
		URL: fmt.Sprintf("http://mypage.aikatsu.com/mypages/index/%s", id),
	}

	res, err := http.Get(idol.URL)
	if err != nil {
		return idol, err
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return idol, err
	}

	// Set Name
	idol.Name = doc.Find("#container > article > div > section > div.m_avatar > div > h4 > span").Text()

	// Set AvatarURL
	au, ok := doc.Find("#container > article > div > section > div.m_avatar > div > img.m_avatar_chara").Attr("src")
	if !ok {
		return idol, fmt.Errorf("Couldn't get AvatarURL")
	}
	idol.AvatarURL = fmt.Sprintf("http://mypage.aikatsu.com%s", au)

	// Set DataGetDate
	re := regexp.MustCompile("データ取得日：(.*)$")
	dgd := re.ReplaceAllString(
		strings.Trim(doc.Find("#container > div.l_header > header > div.m_playdate > p").Text(), " \n"),
		"$1",
	)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	idol.DataGetDate, err = time.ParseInLocation("2006年01月02日 03時04分", dgd, loc)
	if err != nil {
		fmt.Println(err)
	}
	// "2016年01月02日 03時04分"

	// Set LastPlayLocation
	idol.LastPlayLocation = strings.Trim(doc.Find("#container > article > div > section > dl.m_playdate > dd").Text(), " \n")

	// Convert number images to number functions
	// Set each levels
	doc.Find("#container > article > div > section > dl.m_status > dd.m_status_cute > span > img").Each(numberImagesConverterFactory(&idol.CuteLevel))
	doc.Find("#container > article > div > section > dl.m_status > dd.m_status_cool > span > img").Each(numberImagesConverterFactory(&idol.CoolLevel))
	doc.Find("#container > article > div > section > dl.m_status > dd.m_status_sexy > span > img").Each(numberImagesConverterFactory(&idol.SexyLevel))
	doc.Find("#container > article > div > section > dl.m_status > dd.m_status_pop > span > img").Each(numberImagesConverterFactory(&idol.PopLevel))

	// Set misc
	doc.Find("#container > article > div > section > dl.m_rank > dd.m_rank_count > img").Each(numberImagesConverterFactory(&idol.IdolRank))
	idol.IdolRankLabel = strings.Trim(doc.Find("#container > article > div > section > dl.m_rank > dd.m_rank_catch > span").Text(), " \n")
	doc.Find("#container > article > div > section > dl.m_totalfun > dd > span > img").Each(numberImagesConverterFactory(&idol.FanCount))
	doc.Find("#container > article > div > section > dl.m_card > dd > span > img").Each(numberImagesConverterFactory(&idol.PlayedCardsCount))

	return idol, nil
}

// numberImagesConverterFactory creates function that
// converts image element to number and sets value for given counter.
func numberImagesConverterFactory(counter *int) func(int, *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		path := strings.Split(src, "/")
		number, err := strconv.Atoi(path[len(path)-1][:1]) // filename's first char should be number
		if err != nil {
			// error should mean non-number image
			return
		}
		*counter = *counter*10 + number
	}
}
