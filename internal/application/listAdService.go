package application

import (
	"github.com/xejarque/challenges/internal/domain"
	"github.com/xejarque/challenges/internal/infrastructure"
)

type ListAd struct {
	adRepository *infrastructure.InMemoryRepository
}

func NewListAdService(repository *infrastructure.InMemoryRepository) *ListAd {
	return &ListAd{repository}
}

func (s *ListAd) Execute() []domain.Ad {
	return s.adRepository.FindAll()
}
