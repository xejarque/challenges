package domain

import "time"

type Ad struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Created     time.Time
}
