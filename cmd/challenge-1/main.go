package main

import (
	"fmt"
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/application/findById"
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/application/post"
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/ads/infrastructure"
	. "github.com/javier-tw/learning-go/internal/challenge-1/adView/listings/application/create"
)

func main() {
	fmt.Println("Hello world!")
	repository := NewInMemoryRepository()
	postAd := NewPostAd(repository)
	findAdById := NewFindAdById(repository)
	createListing := NewCreateListing(repository)

	postAd.Execute("1")
	err := postAd.Execute("2")
	if err != nil {
		fmt.Printf("We got an error %v", err)
	}

	listing := createListing.Execute("1")
	fmt.Printf("listing is %v\n", listing.Ads)
	ad1 := findAdById.Execute("1")
	fmt.Printf("the ad with id 1 is: %v\n", ad1)
}
