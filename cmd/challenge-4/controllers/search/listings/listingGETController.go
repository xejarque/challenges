package listingsControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"github.com/javier-tw/learning-go/internal/challenge-4/search/listings/application/create"
	"net/http"
)

func ListingGET(repository Ad.Repository) gin.HandlerFunc {
	useCase := listing.NewCreateListing(repository)

	return func(gin *gin.Context) {
		ad, err := useCase.Execute()

		fmt.Printf("%v", ad)
		if err == nil {
			gin.JSON(http.StatusOK, ad)
		} else {
			gin.Status(http.StatusInternalServerError)

		}
		return
	}
}
