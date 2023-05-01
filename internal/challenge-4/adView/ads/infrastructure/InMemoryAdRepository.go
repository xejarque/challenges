package adsInfra

import (
	ads2 "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/application/findById"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"github.com/javier-tw/learning-go/internal/challenge-4/shared/domain/errors"
)

type InMemoryAdsRepository struct{}

var ads []Ad.Ad

func NewInMemoryRepository() InMemoryAdsRepository {
	return InMemoryAdsRepository{}
}
func (InMemoryAdsRepository) SaveAd(ad Ad.Ad) error {
	ads = append(ads, ad)
	return nil
}

func (InMemoryAdsRepository) FindAllAds() ([]Ad.Ad, error) {
	return ads, nil
}

func (InMemoryAdsRepository) FindById(id string) (result Ad.Ad, err error) {
	for _, ad := range ads {
		if ad.Id == id {
			result = ad
			return
		}
	}
	err = domainError.AddValue(ads2.NotFound, id)
	return
}
