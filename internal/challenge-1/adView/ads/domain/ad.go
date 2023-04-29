package ads

import "time"

type Ad struct {
	Id          string
	title       string
	description string
	price       float64
	publishedAt time.Time
}

func CreateAd(id string) *Ad {
	return &Ad{
		Id:          id,
		title:       "Title",
		description: "Such a description",
		price:       213,
		publishedAt: time.Now(),
	}
}
