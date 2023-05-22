package application

import (
	"github.com/xejarque/challenges/internal/domain/ad"
	"github.com/xejarque/challenges/internal/infrastructure"
	"time"
)

type PostAdRequest struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Created     time.Time
}

type PostAd struct {
	AdRepository *infrastructure.InMemoryAdRepository
}

func NewPostAd(repository *infrastructure.InMemoryAdRepository) *PostAd {
	return &PostAd{repository}
}

func (s *PostAd) Execute(request PostAdRequest) {
	err := s.AdRepository.Persist(ad.Ad{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Created:     request.Created,
	})
	if err != nil {
		return
	}
}
