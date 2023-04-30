package ads

import (
	Ad "github.com/javier-tw/learning-go/internal/challenge-3/adView/ads/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindAdById_Should_Return(t *testing.T) {
	repository := NewMockRepository(t)
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	expectedAd, _ := Ad.Create(id, "test description")
	repository.EXPECT().FindById(id).Return(expectedAd, nil)

	useCase := NewFindAdById(repository)
	ad, err := useCase.Execute(id)

	repository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedAd, ad)
}

func Test_FindAdById_Should_Error_When_Is_Not_Found(t *testing.T) {
	repository := NewMockRepository(t)
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"
	repository.EXPECT().FindById(id).Return(Ad.Ad{}, Ad.NotFound)

	useCase := NewFindAdById(repository)
	ad, err := useCase.Execute(id)

	repository.AssertExpectations(t)
	assert.Zero(t, ad)
	assert.ErrorIs(t, err, Ad.NotFound)
}
