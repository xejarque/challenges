package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xejarque/challenges/internal/application"
	"github.com/xejarque/challenges/internal/infrastructure"
	"net/http"
)

func setupRouter() *gin.Engine {
	adRepository := infrastructure.NewInMemoryRepository()
	listAds := application.ListAd{AdRepository: adRepository}

	r := gin.Default()

	r.GET("/api/ads", func(context *gin.Context) {

		context.JSON(http.StatusOK, listAds.Execute())
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
