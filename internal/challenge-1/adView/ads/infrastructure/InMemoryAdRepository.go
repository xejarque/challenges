package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type InMemoryAdsRepository struct{}

var ads []Ad

func NewInMemoryRepository() InMemoryAdsRepository {
	return InMemoryAdsRepository{}
}
func (InMemoryAdsRepository) SaveAd(ad Ad) error {
	ads = append(ads, ad)
	return nil
}

func (InMemoryAdsRepository) FindAllAds() []Ad {
	return ads
}

func (InMemoryAdsRepository) FindById(id string) (result Ad) {
	for _, ad := range ads {
		if ad.Id == id {
			result = ad
		}
	}
	return
	// todo: how to handle not found?
}
