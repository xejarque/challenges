package listing

import (
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/domain"
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/listings/domain"
)

type CreateListingUseCase struct {
	repository Repository
}

func NewCreateListing(repository Repository) CreateListingUseCase {
	return CreateListingUseCase{repository: repository}
}

func (useCase CreateListingUseCase) Execute(id string) (result Listing, err error) {
	ads, err := useCase.repository.FindAllAds()
	if err != nil {
		return
	}
	if len(ads) < 5 {
		result = *CreateListing(id, ads[:])
	} else {
		result = *CreateListing(id, ads[:5])
	}
	return
}
