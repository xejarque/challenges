package main

import (
	"bytes"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectedJson = "[{\"Id\":\"1\",\"Title\":\"title1\",\"Description\":\"description1\",\"Price\":10," +
	"\"Created\":\"<<PRESENCE>>\"},{\"Id\":\"2\",\"Title\":\"title2\",\"Description\":\"description2\"," +
	"\"Price\":20,\"Created\":\"<<PRESENCE>>\"}]"

func TestHandler_ListAds(t *testing.T) {
	ja := jsonassert.New(t)
	router := setupRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/ads", nil)

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	ja.Assertf(recorder.Body.String(), expectedJson)
}

func TestHandler_PostAd(t *testing.T) {
	router := setupRouter()
	recorder := httptest.NewRecorder()
	reqBody := []byte("your request body")
	request, _ := http.NewRequest("PUT", "/api/ads", bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 201, recorder.Code)
	assert.Equal(t, 3, len(AdRepository.Ads))

}
