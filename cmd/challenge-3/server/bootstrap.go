package server

import (
	"github.com/gin-gonic/gin"
	adsControllers "github.com/javier-tw/learning-go/cmd/challenge-3/controllers/adView/ads"
	listingsControllers "github.com/javier-tw/learning-go/cmd/challenge-3/controllers/adView/listings"
	ads2 "github.com/javier-tw/learning-go/internal/challenge-3/adView/ads/infrastructure"
)

func SetupEngine() *gin.Engine {
	engine := gin.Default()

	adRepository := ads2.NewInMemoryRepository()

	engine.PUT("/ads/:id", adsControllers.PostAdPut(adRepository))
	engine.GET("/ads/:id", adsControllers.AdGET(adRepository))
	engine.GET("/listing", listingsControllers.ListingGET(adRepository))
	return engine
}
