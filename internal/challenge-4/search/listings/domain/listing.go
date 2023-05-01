package Listing

import (
	"github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
)

type Listing struct {
	Ads Ad.Ads
}

func Create(ads Ad.Ads) (listing *Listing) {
	if len(ads) < 5 {
		listing = &Listing{Ads: ads[:]}
	} else {
		listing = &Listing{Ads: ads[:5]}
	}
	return
}
