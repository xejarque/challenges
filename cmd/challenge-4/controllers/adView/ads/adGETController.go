package adsControllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	ads "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/application/findById"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"net/http"
)

func AdGET(repository Ad.Repository) gin.HandlerFunc {
	useCase := ads.NewFindAdById(repository)

	return func(gin *gin.Context) {
		id := gin.Param("id")

		ad, err := useCase.Execute(id)

		gin.JSON(http.StatusOK, ad)
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
