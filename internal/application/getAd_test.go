package application

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xejarque/challenges/internal/application/mocks"
	"github.com/xejarque/challenges/internal/domain/ad"
	"testing"
	"time"
)

func Test_Should_Get_Ad_By_Id(t *testing.T) {
	expectedAd := anAd()
	repository := mocks.NewAdRepository(t)
	repository.EXPECT().FindById(mock.Anything).Return(expectedAd, nil)
	getAd := GetAd{repository}

	actualAd, _ := getAd.Execute(GetAdRequest{Id: "id"})

	repository.AssertCalled(t, "FindById", "id")
	assert.Equal(t, expectedAd, actualAd)
}

func Test_Should_Not_Found_Ad_By_Id(t *testing.T) {
	repository := mocks.NewAdRepository(t)
	repository.EXPECT().FindById(mock.Anything).Return(ad.Ad{}, ad.AdNotFound)
	getAd := GetAd{repository}

	_, actualError := getAd.Execute(GetAdRequest{Id: "id"})

	assert.Equal(t, ad.AdNotFound, actualError)
}

func anAd() ad.Ad {
	return ad.Ad{
		Id:          "id",
		Title:       "title",
		Description: "description",
		Price:       10,
		Created:     time.Now(),
	}
}
