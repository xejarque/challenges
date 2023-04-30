package t_test

import (
	"bytes"
	"encoding/json"
	adsControllers "github.com/javier-tw/learning-go/cmd/challenge-3/controllers/adView/ads"
	"github.com/javier-tw/learning-go/cmd/challenge-3/server"
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

	req, _ := http.NewRequest("PUT", "/ads/ef8ac118-8d7f-49cc-abec-78e0d05af80a", bytes.NewBuffer(jsonRequest))
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "", response.Body.String())
}
