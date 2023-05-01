package t_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	adsControllers "github.com/javier-tw/learning-go/cmd/challenge-4/controllers/adView/ads"
	"github.com/javier-tw/learning-go/cmd/challenge-4/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestE2EPostAd(t *testing.T) {
	router := server.SetupEngine()
	response := httptest.NewRecorder()
	request := adsControllers.PostAdRequest{Description: "such a post!"}
	jsonRequest, _ := json.Marshal(request)

	req, _ := http.NewRequest("PUT", "/ads/"+uuid.New().String(), bytes.NewBuffer(jsonRequest))
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "", response.Body.String())
}
