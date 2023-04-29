package listing

import (
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/listings/domain"
)

type CreateListingUseCase struct {
	repository AdRepository
}

func NewCreateListing(repository AdRepository) CreateListingUseCase {
	return CreateListingUseCase{repository: repository}
}

func (useCase CreateListingUseCase) Execute(id string) (result Listing) {
	ads := useCase.repository.FindAllAds()
	if len(ads) < 5 {
		result = *CreateListing(id, ads[:])
	} else {
		result = *CreateListing(id, ads[:5])
	}
	return
}
