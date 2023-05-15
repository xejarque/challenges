package application

import (
	"github.com/xejarque/challenges/internal/domain"
	"github.com/xejarque/challenges/internal/infrastructure"
)

type GetAdRequest struct {
	Id string
}

type GetAd struct {
	adRepository *infrastructure.InMemoryRepository
}

func NewGetAd(repository *infrastructure.InMemoryRepository) *GetAd {
	return &GetAd{repository}
}

func (s *GetAd) Execute(request GetAdRequest) domain.Ad {
	return s.adRepository.FindByAdId(request.Id)
}
