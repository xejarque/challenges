package main

import (
	"fmt"
	"github.com/xejarque/challenges/internal/application"
	"github.com/xejarque/challenges/internal/infrastructure"
	"time"
)

func main() {
	var inMemoryRepository = infrastructure.NewInMemoryRepository()
	var service = application.NewPostAd(inMemoryRepository)
	service.PostAd(application.PostingAdRequest{
		Id:          "UUID3",
		Title:       "title",
		Description: "description",
		Price:       20,
		Created:     time.Now(),
	})

	fmt.Println(inMemoryRepository.Ads)
}
