package Listing

import (
	"github.com/javier-tw/learning-go/internal/challenge-3/adView/ads/domain"
)

type Listing struct {
	Ads []Ad.Ad
}

func Create(ads []Ad.Ad) *Listing {
	return &Listing{Ads: ads}
}
