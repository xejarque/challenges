package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type PostAd struct {
	repository AdRepository
}

func NewPostAd(repository AdRepository) PostAd {
	return PostAd{repository: repository}
}

func (useCase PostAd) Execute(id string) error {
	ad := CreateAd(id)
	err := useCase.repository.SaveAd(*ad)
	if err != nil {
		return err
	}
	return nil
}
