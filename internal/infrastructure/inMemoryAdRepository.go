package infrastructure

import (
	. "github.com/xejarque/challenges/internal/domain/ad"
	"time"
)

type InMemoryAdRepository struct {
	Ads map[string]Ad
}

func NewInMemoryRepository() *InMemoryAdRepository {
	adRepository := &InMemoryAdRepository{
		make(map[string]Ad),
	}
	adRepository.Ads["UUID1"] = Ad{
		Id:          "1",
		Title:       "title1",
		Description: "description1",
		Price:       10,
		Created:     time.Now(),
	}
	adRepository.Ads["UUID2"] = Ad{
		Id:          "2",
		Title:       "title2",
		Description: "description2",
		Price:       20,
		Created:     time.Now(),
	}
	return adRepository
}

func (r *InMemoryAdRepository) FindAll() ([]Ad, error) {
	return getMapValues(r.Ads), nil
}

func (r *InMemoryAdRepository) Persist(ad Ad) error {
	ad = Ad{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Created:     ad.Created,
	}
	r.Ads[ad.Id] = ad
	return nil
}
func (r *InMemoryAdRepository) FindById(adId string) (result *Ad, err error) {
	if adFound, ok := r.Ads[adId]; ok {
		result = &adFound
	} else {
		err = AddError(AdNotFound, adId)
		result = nil
	}
	return
}

func getMapValues(m map[string]Ad) []Ad {
	values := make([]Ad, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
