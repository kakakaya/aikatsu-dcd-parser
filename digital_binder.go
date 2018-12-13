package dcdkatsu

import (
	"fmt"
	"net/http"
	"regexp"
	// 	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

// DigitalBinder holds acquisition status of series of cards.
// Series means binder's season.
type DigitalBinder struct {
	SeriesID   int    `json:"series_id"`
	SeriesName string `json:"series_name"`

	Cards []Card `json:"cards"`

	// Owner   Idol   `json:"-"`
	OwnerID string `json:"owner_id"`
	URL     string `json:"url"`

	DataGetDate time.Time `json:"data_get_date"`
}

// FetchDigitalBinder returns parsed DigitalBinder data.
// This data will be fetched from http://mypage.aikatsu.com/mypages/digital_binders/<ownerID>/<seriesID> .
//
// Note: Card's detail won't be fetchedf, Because each of them requires one HTTP request per card.
// Use Card.Detail() for more information.
func FetchDigitalBinder(ownerID string, seriesID int) (DigitalBinder, error) {
	db := DigitalBinder{
		OwnerID:  ownerID,
		SeriesID: seriesID,
		URL:      fmt.Sprintf("http://mypage.aikatsu.com/mypages/digital_binders/%s/%d", ownerID, seriesID),
	}

	res, err := http.Get(db.URL)
	if err != nil {
		return db, err
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return db, err
	}

	db.SeriesName = strings.Trim(doc.Find("#container > article > div > section > div.l_box > div.m_tit > span").Text(), " \n")

	doc.Find("#container > article > div > section > div.l_box > div.m_inner > ul.m_dress > li").Each(parseAndSetCard(&db))

	// Set DataGetDate
	db.DataGetDate, err = parseDataGetDate(strings.Trim(doc.Find("#container > div.l_header > header > div.m_playdate > p").Text(), " \n"))
	if err != nil {
		return db, err
	}

	return db, nil
}

// parseAndSetCard creats function that
// parses card info from binder.
func parseAndSetCard(db *DigitalBinder) func(i int, s *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		card := Card{
			OwnerID:  db.OwnerID,
			SeriesID: db.SeriesID,
		}

		re := regexp.MustCompile(`^.*?\('(.*?)'\).*?$`)
		path, _ := s.Find("a").Attr("onclick")
		card.URL = re.ReplaceAllString(path, "http://mypage.aikatsu.com$1")

		src, _ := s.Find("a > div.m_dress_card_img > img").First().Attr("src")
		card.ImageURL = fmt.Sprintf("http://mypage.aikatsu.com/%s", src)

		// Set owned
		if s.Find("a > div.m_dress_card_img > img.is_medal").Length() == 0 {
			card.Owned = false
		} else {
			card.Owned = true
		}

		db.Cards = append(db.Cards, card)
	}
}
