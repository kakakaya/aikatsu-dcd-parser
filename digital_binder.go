package dcdkatsu

import (
	// 	"fmt"
	// 	"net/http"
	// 	"regexp"
	// 	"strconv"
	// 	"strings"
	"time"
	// 	"github.com/PuerkitoBio/goquery"
)

// DigitalBinder holds acquisition status of series of cards.
// Series means binder's season.
type DigitalBinder struct {
	SeriesID   int    `json:"series_id"`
	SeriesName string `json:"series_name"`

	Cards []Card `json:"-"`

	Owner   Idol   `json:"-"`
	OwnerID string `json:"owner_id"`
	URL     string `json:"url"`

	DataGetDate time.Time `json:"data_get_date"`
}
