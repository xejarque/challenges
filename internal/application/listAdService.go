package application

import (
	"github.com/xejarque/challenges/internal/domain/ad"
	"github.com/xejarque/challenges/internal/infrastructure"
)

type ListAd struct {
	adRepository *infrastructure.InMemoryAdRepository
}

func NewListAdService(repository *infrastructure.InMemoryAdRepository) *ListAd {
	return &ListAd{repository}
}

func (s *ListAd) Execute() []ad.Ad {
	all, _ := s.adRepository.FindAll()
	return all
}
