package api

import (
	"fmt"
	"net/http"

	"github.com/andreibthoughtworks/challenges/internal/domain"
	"github.com/gin-gonic/gin"
)

func putAd(ctx *gin.Context, store domain.AdStore) {
	ad := &domain.AdRequest{}
	err := ctx.BindJSON(ad)
	if err != nil {
		fmt.Printf("invalid json: %v\n", err)

		return
	}

	storedAd, err := store.Store(*ad)
	if err != nil {
		ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf("could not store ad: %v", err),
		)

		return
	}

	ctx.JSON(http.StatusOK, &storedAd)
}

func getAd(ctx *gin.Context, store domain.AdStore) {
	id := ctx.Param("id")
	ad, err := store.Get(id)
	if err != nil {
		ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf("could not find ad '%s'': %v", id, err),
		)

		return
	}

	if ad == nil {
		ctx.String(http.StatusNotFound,
			fmt.Sprintf("could not find ad '%s'", id))

		return
	}

	ctx.JSON(http.StatusOK, ad)
}

func getAds(ctx *gin.Context, store domain.AdStore) {
	ads, err := store.GetMany()
	if err != nil {
		ctx.String(
			http.StatusInternalServerError,
			fmt.Sprintf("could not get ads: %v", err),
		)
	}

	ctx.JSON(http.StatusOK, ads)
}

func StartServer(store domain.AdStore) {
	router := gin.Default()

	// closure in the ad store to conform to interface
	putAdHandler := func(c *gin.Context) {
		putAd(c, store)
	}
	router.PUT("/ads", putAdHandler)

	getAdHandler := func(c *gin.Context) {
		getAd(c, store)
	}
	router.GET("/ads/:id", getAdHandler)

	getAdsHandler := func(c *gin.Context) {
		getAds(c, store)
	}
	router.GET("/ads", getAdsHandler)

	router.Run("localhost:8080")
}
