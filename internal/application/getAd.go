package application

import (
	"github.com/xejarque/challenges/internal/domain/ad"
)

type GetAdRequest struct {
	Id string
}

type GetAd struct {
	adRepository AdRepository
}

func NewGetAd(repository AdRepository) *GetAd {
	return &GetAd{repository}
}

func (s *GetAd) Execute(request GetAdRequest) (ad.Ad, error) {
	return s.adRepository.FindById(request.Id)
}
