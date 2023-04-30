package main

import (
	"errors"
	"fmt"
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/application/findById"
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/application/post"
	ad "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/domain"
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/ads/infrastructure"
	. "github.com/javier-tw/learning-go/internal/challenge-2/adView/listings/application/create"
)

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("|Hello world from challenge - 2!|")
	fmt.Println("---------------------------------")

	repository := NewInMemoryRepository()
	postAd := NewPostAd(repository)
	findAdById := NewFindAdById(repository)
	createListing := NewCreateListing(repository)

	err := postAd.Execute("1", "this ad should be fine")
	if err != nil {
		handleError(err)
	}

	err = postAd.Execute("2", "This ad is to loong and should throw and errror because is longer than 50 characters")
	if err != nil {
		if errors.Is(err, ad.InvalidAdDescriptionError) {
			handleError(err)
		}
	}

	listing, err := createListing.Execute("1")
	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("listing is %v\n", listing.Ads)
	}

	ad1, err := findAdById.Execute("1")
	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("the ad with id 1 is: %v\n", ad1)
	}

	ad2, err := findAdById.Execute("2")
	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("the ad with id 1 is: %v\n", ad2)
	}
}

func handleError(err error) {
	fmt.Printf("%v\n", err)
}
