package dcd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Card represents a card.
// Any card would be fetched with user id, so Card has user's ID as "OwnerID".
type Card struct {
	Code string `json:"code"`
	Name string `json:"name"`

	Rarity string `json:"rarity"`
	Stars  string `json:"stars"`

	Type        string `json:"type"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	DressAppeal string `json:"dress_appeal"`
	AppealPoint int    `json:"appeal_point"`

	Owner   Idol   `json:"-"`
	OwnerID string `json:"owner_id"`
	Owned   bool   `json:"owned"`

	SeriesID int    `json:"series_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`

	detailed bool // true after Card.Detail called and succeed.
}

// FetchCard returns parsed Idol data for given url.
func FetchCard(url string) (Card, error) {
	var err error

	// trim trailing slash
	if string(url[len(url)-1]) == "/" {
		url = url[0 : len(url)-1]
	}
	card := Card{
		URL: url,
	}
	p := strings.Split(url, "/")
	card.OwnerID = p[len(p)-1]
	card.SeriesID, err = strconv.Atoi(p[len(p)-2])
	if err != nil {
		return card, err
	}

	res, err := http.Get(card.URL)
	if err != nil {
		return card, err
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return card, err
	}

	header := strings.Split(strings.Trim(doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > h5").Text(), " \n"), " ")
	if len(header) != 3 {
		return card, fmt.Errorf("Card parsing error: %v", header)
	}
	card.Code = header[0]
	card.Name = header[2]

	card.Type = doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_type > dd > img").AttrOr("alt", "")
	card.Category = doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_attribute > dd > span").Text()

	_brand := doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_brand > dd > span")
	if _brand.Length() == 0 {
		card.Brand = doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_brand > dd > img").AttrOr("alt", "")
	} else {
		card.Brand = _brand.Text()
	}

	card.DressAppeal = doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_appeal > dd > span").Text()
	card.AppealPoint, err = strconv.Atoi(doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_point > dd > span").Text())
	if err != nil {
		return card, err
	}
	// Set owned
	if doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_imgCol > div > img.is_medal").Length() == 0 {
		card.Owned = false
	} else {
		card.Owned = true
	}
	card.ImageURL = fmt.Sprintf("http://mypage.aikatsu.com%s", doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_imgCol > div > img").First().AttrOr("src", ""))

	_rare := doc.Find("body > div.l_dress-datail > div > div.m_dress-datail-box > div > div.m_dataCol > dl.m_rare > dd > span").Text()
	card.Stars = strings.Repeat("★", strings.Count(_rare, "★"))
	card.Rarity = strings.Trim(_rare, "★")

	return card, nil
}

// Detail sets and .
func (card *Card) Detail() (Card, error) {
	if card.detailed {
		return *card, nil
	}
	fmt.Println(card.ImageURL)
	card.detailed = true
	return *card, nil
}
