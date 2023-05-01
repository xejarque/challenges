package adsControllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	ads2 "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/application/post"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"net/http"
)

type PostAdRequest struct {
	Description string `json:"description" binding:"required"`
}

func PostAdPut(repository Ad.Repository) gin.HandlerFunc {
	useCase := ads2.NewPostAd(repository)

	return func(gin *gin.Context) {
		id := gin.Param("id")
		var request PostAdRequest
		if err := gin.BindJSON(&request); err != nil {
			gin.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := useCase.Execute(id, request.Description)

		gin.Status(http.StatusCreated)
		if err != nil {
			if errors.Is(err, Ad.InvalidAdDescriptionError) {
				gin.JSON(http.StatusBadRequest, err.Error())
				return
			} else {
				gin.Status(http.StatusInternalServerError)
				return
			}
		}
		return
	}
}
