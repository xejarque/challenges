package ad

import "time"

var AdNotFound = CreateDomainError("the ad was not found")

type Ad struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Created     time.Time
}
