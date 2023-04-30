package Ad

type Repository interface {
	SaveAd(ad Ad) error
	FindAllAds() []Ad
	FindById(id string) Ad
}
