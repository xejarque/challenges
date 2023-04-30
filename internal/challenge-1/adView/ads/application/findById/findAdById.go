package ads

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type FindAdById struct {
	repository Ad.Repository
}

func NewFindAdById(repository Ad.Repository) FindAdById {
	return FindAdById{repository: repository}
}

func (useCase FindAdById) Execute(id string) Ad.Ad {
	return useCase.repository.FindById(id)
}
