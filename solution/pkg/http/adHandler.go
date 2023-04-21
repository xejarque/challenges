package http

// TODO: I figured http handlers should probably go here, still to-do

//
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//func putAd(ctx *gin.Context, store *ads.AdStore) {
//	ad := &ads.Ad{}
//	err := ctx.BindJSON(ad)
//	if err != nil {
//		fmt.Printf("invalid json: %v\n", err)
//
//		return
//	}
//
//	if valid, reason := ad.IsValid(); !valid {
//		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid ad: %s", reason)).
//			SetType(gin.ErrorTypePublic)
//
//		return
//	}
//
//	storedAd, err := store.Store(ad.Title, ad.Description, ad.Price)
//	if err != nil {
//		ctx.String(
//			http.StatusInternalServerError,
//			fmt.Sprintf("could not store ad: %v", err),
//		)
//
//		return
//	}
//
//	ctx.JSON(http.StatusOK, &storedAd)
//}
//
//func getAd(ctx *gin.Context, store *ads.AdStore) {
//	id := ctx.Param("id")
//	ad, err := store.Find(id)
//	if err != nil {
//		ctx.String(
//			http.StatusInternalServerError,
//			fmt.Sprintf("could not find ad '%s'': %v", id, err),
//		)
//
//		return
//	}
//
//	if ad == nil {
//		ctx.String(http.StatusNotFound,
//			fmt.Sprintf("could not find ad '%s'", id))
//
//		return
//	}
//
//	ctx.JSON(http.StatusOK, ad)
//}
