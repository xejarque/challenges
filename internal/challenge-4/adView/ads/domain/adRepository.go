package Ad

//go:generate mockery --name Repository
type Repository interface {
	SaveAd(ad Ad) error
	FindAllAds() (Ads, error)
	FindById(id string) (Ad, error)
}
