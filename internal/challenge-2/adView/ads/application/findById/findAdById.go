package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/domain"
)

type FindAdById struct {
	repository Repository
}

func NewFindAdById(repository Repository) FindAdById {
	return FindAdById{repository: repository}
}

func (useCase FindAdById) Execute(id string) (Ad, error) {
	return useCase.repository.FindById(id)
}
