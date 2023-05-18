package application

import (
	. "github.com/xejarque/challenges/internal/domain/ad"
)

//go:generate mockery --name AdRepository
type AdRepository interface {
	Persist(ad Ad) error
	FindAll() ([]Ad, error)
	FindById(id string) (Ad, error)
}
