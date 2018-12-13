package dcdkatsu

import (
	"fmt"
	"strings"
)

// Card represents a card.
// Any card would be fetched with user id, so Card has user's ID as "OwnerID".
type Card struct {
	Code   string `json:"code"`
	Rarity string `json:"rarity"`
	Name   string `json:"name"`

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

// FetchCard returns parsed Idol data.
func FetchCard(url string) (Card, error) {
	card := Card{
		URL: url,
	}
	p := strings.Split(url, "/")
	card.OwnerID = p[len(p)-2]
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
