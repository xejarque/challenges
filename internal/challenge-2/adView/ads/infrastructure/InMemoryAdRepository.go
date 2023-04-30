package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/domain"
	domainError "github.com/javier-tw/learning-go/internal/challenge-2/adView/shared/domain"
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

func (InMemoryAdsRepository) FindAllAds() ([]Ad, error) {
	return ads, nil
}

func (InMemoryAdsRepository) FindById(id string) (result Ad, err error) {
	for _, ad := range ads {
		if ad.Id == id {
			result = ad
			return
		}
	}
	err = domainError.AddValue(NotFound, id)
	return
}
