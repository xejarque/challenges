package listing

import (
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type Listing struct {
	id  string
	Ads []Ad
}

func CreateListing(id string, ads []Ad) *Listing {
	return &Listing{id: id, Ads: ads}
}
