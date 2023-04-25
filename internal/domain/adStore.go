package domain

import (
	"fmt"

	"github.com/andreibthoughtworks/challenges/pkg/util"
)

type AdStore interface {
	Store(adReq AdRequest) (*ad, error)
	Get(id string) (*ad, error)
	GetMany() ([]ad, error)
}

type MemoryStore struct {
	store map[string]ad
}

func NewMemoryStore() AdStore {
	return &MemoryStore{
		store: map[string]ad{},
	}
}

func (s *MemoryStore) Store(adReq AdRequest) (*ad, error) {
	ad, invalidReason := newAd(adReq)
	if ad == nil {
		return nil, fmt.Errorf("invalid ad: %s", *invalidReason)
	}

	s.store[ad.ID] = *ad

	return ad, nil
}

func (s *MemoryStore) Get(id string) (*ad, error) {
	ad, found := s.store[id]
	if !found {
		return nil, nil
	}

	return util.GetPointer(ad), nil
}

func (s *MemoryStore) GetMany() ([]ad, error) {
	end := 5
	if len(s.store) < 5 {
		end = len(s.store)
	}

	ads := make([]ad, end)
	i := 0
	for _, a := range s.store {
		ads[i] = a
		i++
		if i >= end {
			break
		}
	}

	return ads, nil
}
