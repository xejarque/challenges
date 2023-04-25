package main

import (
	"github.com/andreibthoughtworks/challenges/internal/api"
	"github.com/andreibthoughtworks/challenges/internal/domain"
)

func main() {
	store := domain.NewMemoryStore()

	api.StartServer(store)
}
