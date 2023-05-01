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

type Primitives struct {
	Id          string
	Title       string
	Description string
	Price       float64
	PublishedAt time.Time
}

func Create(id string, description string) (Ad, error) {
	adDescription, err := CreateDescription(description)
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

func (a Ad) ToPrimitives() Primitives {
	return Primitives{
		Id:          a.Id,
		Title:       a.Title,
		Description: a.Description.String(),
		Price:       a.Price,
		PublishedAt: a.PublishedAt,
	}
}

func FromPrimitives(primitives Primitives) Ad {
	return Ad{
		Id:          primitives.Id,
		Title:       primitives.Title,
		Description: DescriptionFromPrimitives(primitives.Description),
		Price:       primitives.Price,
		PublishedAt: primitives.PublishedAt,
	}
}
