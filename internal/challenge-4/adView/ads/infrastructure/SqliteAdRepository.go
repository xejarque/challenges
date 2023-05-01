package adsInfra

import (
	"errors"
	Ad "github.com/javier-tw/learning-go/internal/challenge-4/adView/ads/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type SqliteAdsRepository struct {
	client *gorm.DB
}

type AdSchema struct {
	Id          string `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float64
	PublishedAt time.Time
}

func NewSqliteRepository() SqliteAdsRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&AdSchema{})
	if err != nil {
		panic("failed to connect database")
	}
	return SqliteAdsRepository{client: db}
}

func (r SqliteAdsRepository) SaveAd(ad Ad.Ad) error {
	return r.client.Create(ad.ToPrimitives()).Error
}

func (r SqliteAdsRepository) FindAllAds() (Ad.Ads, error) {
	var adsPrimitives []Ad.Primitives
	if err := r.client.Find(&adsPrimitives).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Ad.Ad{}, nil
		}
		return []Ad.Ad{}, err
	}
	return Ad.CollectionFromPrimitives(adsPrimitives), nil
}

func (r SqliteAdsRepository) FindById(id string) (Ad.Ad, error) {
	var adPrimitives Ad.Primitives

	if err := r.client.First(&adPrimitives, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Ad.Ad{}, nil
		}
		return Ad.Ad{}, err
	}
	return Ad.FromPrimitives(adPrimitives), nil
}
