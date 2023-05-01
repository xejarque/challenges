package server

import (
	"github.com/gin-gonic/gin"
	adsControllers "github.com/javier-tw/learning-go/cmd/challenge-4/controllers/adView/ads"
	"github.com/javier-tw/learning-go/cmd/challenge-4/controllers/search/listings"
	adsInfra "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/infrastructure"
)

func SetupEngine() *gin.Engine {
	engine := gin.Default()

	adRepository := adsInfra.NewSqliteRepository()

	engine.PUT("/ads/:id", adsControllers.PostAdPut(adRepository))
	engine.GET("/ads/:id", adsControllers.AdGET(adRepository))
	engine.GET("/listing", listingsControllers.ListingGET(adRepository))
	return engine
}
