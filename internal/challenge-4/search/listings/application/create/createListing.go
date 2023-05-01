package listing

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"github.com/javier-tw/learning-go/internal/challenge-4/search/listings/domain"
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
	result = *Listing.Create(ads)
	return
}
