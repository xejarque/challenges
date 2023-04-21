package domain

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	ErrorAdNotFound = errors.New("ad not found")
)

type AdID string

type (
	AdRepository interface {
		Create(context.Context, Ad) (Ad, error)
		FindAll(context.Context) ([]Ad, error)
		FindByID(context.Context, string) (Ad, error)
	}

	Ad struct {
		ID          string
		Title       string
		Description string
		Price       int
		CreatedAt   time.Time
	}
)

func NewAd(title, description string, price int) Ad {
	return Ad{
		ID:          uuid.NewV4().String(),
		Title:       title,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
	}
}
