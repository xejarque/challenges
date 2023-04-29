package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/domain"
)

type FindAdById struct {
	repository AdRepository
}

func NewFindAdById(repository AdRepository) FindAdById {
	return FindAdById{repository: repository}
}

func (useCase FindAdById) Execute(id string) Ad {
	return useCase.repository.FindById(id)
}
