package t_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	adsControllers "github.com/javier-tw/learning-go/cmd/challenge-4/controllers/adView/ads"
	"github.com/javier-tw/learning-go/cmd/challenge-4/server"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestE2EGetAd(t *testing.T) {
	engine := server.SetupEngine()
	expectedAd := PostAd(engine)
	response := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ads/"+expectedAd.Id, nil)
	engine.ServeHTTP(response, req)

	var result Ad.Ad
	json.Unmarshal(response.Body.Bytes(), &result)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, expectedAd.Id, result.Id)
	assert.Equal(t, expectedAd.Description, result.Description)
}

func PostAd(engine *gin.Engine) Ad.Ad {
	id := uuid.New().String()
	description, _ := Ad.CreateDescription("such a post!")
	response := httptest.NewRecorder()
	request := adsControllers.PostAdRequest{Description: description.String()}
	jsonRequest, _ := json.Marshal(request)
	req, _ := http.NewRequest("PUT", "/ads/"+id, bytes.NewBuffer(jsonRequest))
	engine.ServeHTTP(response, req)
	return Ad.Ad{Id: id, Description: description, Title: "Title", Price: 213, PublishedAt: time.Now()}
}
