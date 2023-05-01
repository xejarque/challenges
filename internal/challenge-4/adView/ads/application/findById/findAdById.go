package ads

import (
	. "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"github.com/javier-tw/learning-go/internal/challenge-4/shared/domain/errors"
)

type FindAdById struct {
	repository Repository
}

func NewFindAdById(repository Repository) FindAdById {
	return FindAdById{repository: repository}
}

func (useCase FindAdById) Execute(id string) (Ad, error) {
	ad, err := useCase.repository.FindById(id)
	if ad.Id == "" {
		return Ad{}, domainError.AddValue(NotFound, id)
	}
	return ad, err
}
