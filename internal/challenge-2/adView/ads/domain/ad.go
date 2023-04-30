package Ad

import (
	"time"
)

type Ad struct {
	Id          string
	title       string
	description Description
	price       float64
	publishedAt time.Time
}

func Create(id string, description string) (Ad, error) {
	adDescription, err := CreateDescription(description)
	if err != nil {
		return Ad{}, err
	}
	return Ad{
		Id:          id,
		title:       "Title",
		description: adDescription,
		price:       213,
		publishedAt: time.Now(),
	}, nil
}
