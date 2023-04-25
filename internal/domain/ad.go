package domain

import (
	"time"

	"github.com/andreibthoughtworks/challenges/pkg/util"
	"github.com/google/uuid"
)

type InvalidReason string

const DescTooLong InvalidReason = "The Description must be 50 characters or less"

type AdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ad struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	AdRequest
}

func newAd(adReq AdRequest) (*ad, *InvalidReason) {
	ad := ad{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		AdRequest: adReq,
	}

	if valid, reason := ad.isValid(); !valid {
		return nil, reason
	}

	return &ad, nil
}

func (a *ad) isValid() (bool, *InvalidReason) {
	if len(a.Description) > 50 {
		return false, util.GetPointer(DescTooLong)
	}

	return true, nil
}
