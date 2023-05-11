package application

import (
	"github.com/xejarque/challenges/internal/domain"
	"github.com/xejarque/challenges/internal/infrastructure"
	"time"
)

type PostingAdRequest struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Created     time.Time
}

type PostAd struct {
	adRepository *infrastructure.InMemoryRepository
}

func NewPostAd(repository *infrastructure.InMemoryRepository) *PostAd {
	return &PostAd{repository}
}

func (s *PostAd) PostAd(request PostingAdRequest) {
	s.adRepository.Persist(domain.Ad{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Created:     request.Created,
	})
}
