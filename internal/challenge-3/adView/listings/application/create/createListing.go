package listing

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-3/adView/ads/domain"
	Listing "github.com/javier-tw/learning-go/internal/challenge-3/adView/listings/domain"
)

type CreateListingUseCase struct {
	repository Ad.Repository
}

func NewCreateListing(repository Ad.Repository) CreateListingUseCase {
	return CreateListingUseCase{repository: repository}
}

func (useCase CreateListingUseCase) Execute() (result Listing.Listing, err error) {
	ads, err := useCase.repository.FindAllAds()
	if err != nil {
		return
	}
	if len(ads) < 5 {
		result = *Listing.Create(ads[:])
	} else {
		result = *Listing.Create(ads[:5])
	}
	return
}
