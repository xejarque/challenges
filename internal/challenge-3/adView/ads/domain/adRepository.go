package Ad

//go:generate mockery --name Repository
type Repository interface {
	SaveAd(ad Ad) error
	FindAllAds() ([]Ad, error)
	FindById(id string) (Ad, error)
}
