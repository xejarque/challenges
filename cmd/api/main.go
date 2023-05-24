package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xejarque/challenges/internal/application"
	"github.com/xejarque/challenges/internal/domain/ad"
	"github.com/xejarque/challenges/internal/infrastructure"
	"net/http"
)

var AdRepository = infrastructure.NewInMemoryRepository()

func setupRouter() *gin.Engine {
	listAds := application.NewListAdService(AdRepository)
	postAd := application.NewPostAd(AdRepository)
	getAd := application.NewGetAd(AdRepository)

	engine := gin.Default()
	engine.GET("/api/ads", findAllHandler(listAds))
	engine.GET("/api/ads/:id", getAdHandler(getAd))
	engine.PUT("/api/ads/:id", postAdHandler(postAd))

	return engine
}

func postAdHandler(postAd *application.PostAd) func(context *gin.Context) {
	return func(context *gin.Context) {
		var postAdDto struct {
			Id          string  `json:"id" binding:"required"`
			Title       string  `json:"title" binding:"required"`
			Description string  `json:"description" binding:"required"`
			Price       float64 `json:"price" binding:"required"`
		}
		if err := context.BindJSON(&postAdDto); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		response := postAd.Execute(application.PostAdRequest{
			Id:          postAdDto.Id,
			Title:       postAdDto.Title,
			Description: postAdDto.Description,
			Price:       postAdDto.Price,
		})
		context.JSON(http.StatusCreated, response)
	}
}

func getAdHandler(getAd *application.GetAd) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		adFound, err := getAd.Execute(application.GetAdRequest{Id: id})
		if errors.Is(err, ad.AdNotFound) {
			context.JSON(http.StatusNotFound, err.Error())
		}
		context.JSON(http.StatusOK, adFound)
	}
}

func findAllHandler(listAds *application.ListAd) func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, listAds.Execute())
	}
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
