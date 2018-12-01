package dcdkatsu

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Idol hold information for Aikatsu! idol.
type Idol struct {
	ID  string
	URL string

	// Basic Info
	Name      string
	AvatarURL string

	// Levels for each type
	CuteLevel int
	CoolLevel int
	SexyLevel int
	PopLevel  int

	FanCount int

	LastPlayLocation string

	// TODO: Check DailyHighScore type; is it int?
	// If there's no rank, set to zero.
	DailyHighScoreSolo    int
	DailyHighScoreFriends int
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

	// Set LastPlayLocation
	idol.LastPlayLocation = strings.Trim(doc.Find("#container > article > div > section > dl.m_playdate > dd").Text(), " \n")

	// Convert number images to number functions
	fanNumberImagesConverter := numberImagesConverterFactory(&idol.FanCount)
	doc.Find("#container > article > div > section > dl.m_totalfun > dd > span > img").Each(fanNumberImagesConverter)

	return idol, nil
}

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
