package infrastructure

import (
	"fmt"
	"github.com/xejarque/challenges/internal/domain"
	"time"
)

type InMemoryRepository struct {
	Ads map[string]domain.Ad
}

func NewInMemoryRepository() *InMemoryRepository {
	adRepository := &InMemoryRepository{
		make(map[string]domain.Ad),
	}
	adRepository.Ads["UUID1"] = domain.Ad{
		Id:          "1",
		Title:       "title1",
		Description: "description1",
		Price:       10,
		Created:     time.Now(),
	}
	adRepository.Ads["UUID2"] = domain.Ad{
		Id:          "2",
		Title:       "title2",
		Description: "description2",
		Price:       20,
		Created:     time.Now(),
	}
	return adRepository
}

func (r *InMemoryRepository) FindAll() []domain.Ad {
	return getMapValues(r.Ads)
}

func (r *InMemoryRepository) Persist(ad domain.Ad) {
	r.Ads[ad.Id] = domain.Ad{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Created:     ad.Created,
	}
}
func (r *InMemoryRepository) FindByAdId(adId string) domain.Ad {
	fmt.Println("findByAdId repo call")
	return r.Ads[adId]
}

func getMapValues(m map[string]domain.Ad) []domain.Ad {
	values := make([]domain.Ad, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
