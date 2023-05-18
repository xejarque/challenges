package main

import (
	"fmt"
	"github.com/xejarque/challenges/internal/application"
	"github.com/xejarque/challenges/internal/infrastructure"
	"time"
)

func main() {
	//INIT
	var inMemoryRepository = infrastructure.NewInMemoryRepository()
	fmt.Println("######### INIT ##########")
	fmt.Println(inMemoryRepository.Ads)

	//POST AD
	var postAdService = application.NewPostAd(inMemoryRepository)
	postAdService.Execute(application.PostAdRequest{
		Id:          "UUID3",
		Title:       "title",
		Description: "description",
		Price:       20,
		Created:     time.Now(),
	})
	fmt.Println("######### PERSIST NEW AD ##########")
	fmt.Println(inMemoryRepository.Ads)

	//LIST ADS
	fmt.Println("######### LIST ADS ##########")
	var listAdService = application.NewListAdService(inMemoryRepository)
	fmt.Println(listAdService.Execute())

	//FIND AD BY ID
	fmt.Println("######### FIND AD BY ID ##########")
	var getAdService = application.NewGetAd(inMemoryRepository)
	fmt.Println(getAdService.Execute(application.GetAdRequest{Id: "UUID3"}))

}
