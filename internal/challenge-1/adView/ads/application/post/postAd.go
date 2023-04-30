package ads

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type PostAd struct {
	repository Ad.Repository
}

func NewPostAd(repository Ad.Repository) PostAd {
	return PostAd{repository: repository}
}

func (useCase PostAd) Execute(id string) error {
	ad := Ad.Create(id)
	err := useCase.repository.SaveAd(*ad)
	if err != nil {
		return err
	}
	return nil
}
