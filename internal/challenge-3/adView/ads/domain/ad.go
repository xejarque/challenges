package Ad

import (
	"time"
)

type Ad struct {
	Id          string
	Title       string
	Description Description
	Price       float64
	PublishedAt time.Time
}

func Create(id string, description string) (Ad, error) {
	adDescription, err := CreateAdDescription(description)
	if err != nil {
		return Ad{}, err
	}
	return Ad{
		Id:          id,
		Title:       "Title",
		Description: adDescription,
		Price:       213,
		PublishedAt: time.Now(),
	}, nil
}
