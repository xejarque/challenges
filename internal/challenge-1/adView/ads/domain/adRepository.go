package ads

type AdRepository interface {
	SaveAd(ad Ad) error
	FindAllAds() []Ad
	FindById(id string) Ad
}
