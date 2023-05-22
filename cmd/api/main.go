package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xejarque/challenges/internal/application"
	"github.com/xejarque/challenges/internal/infrastructure"
	"net/http"
	"time"
)

var AdRepository = infrastructure.NewInMemoryRepository()

func setupRouter() *gin.Engine {
	listAds := application.ListAd{AdRepository: AdRepository}
	postAd := application.PostAd{AdRepository: AdRepository}

	r := gin.Default()
	r.GET("/api/ads", findAllHandler(listAds))
	r.PUT("/api/ads", postAdHandler(postAd))

	return r
}

func postAdHandler(postAd application.PostAd) func(context *gin.Context) {
	var postAdDto struct {
		Id          string    `json:"id" binding:"required"`
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description" binding:"required"`
		Price       float64   `json:"price" binding:"required"`
		Created     time.Time `json:"created" binding:"required"`
	}
	return func(context *gin.Context) {
		postAd.Execute(application.PostAdRequest{
			Id:          postAdDto.Id,
			Title:       postAdDto.Title,
			Description: postAdDto.Description,
			Price:       postAdDto.Price,
			Created:     postAdDto.Created,
		})
		context.JSON(http.StatusCreated, "")

	}
}

func findAllHandler(listAds application.ListAd) func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, listAds.Execute())
	}
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
