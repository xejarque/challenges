package ads

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-3/adView/ads/domain"
)

type PostAd struct {
	repository Ad.Repository
}

func NewPostAd(repository Ad.Repository) PostAd {
	return PostAd{repository: repository}
}

func (useCase PostAd) Execute(id string, description string) error {
	ad, err := Ad.Create(id, description)
	if err != nil {
		return err
	}
	err = useCase.repository.SaveAd(ad)
	if err != nil {
		return err
	}
	return nil
}
