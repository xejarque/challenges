package app

import (
	"github.com/andreibthoughtworks/challenges/solution/internal/domain"
	"github.com/andreibthoughtworks/challenges/solution/internal/infrastructure"
)

type Application struct {
	adRepository domain.AdRepository
	dbAdapter    infrastructure.SqlDB
}

func NewApplication(dbAdapter *infrastructure.SqlDB) *Application {
	// Create a new instance of the application with the necessary dependencies
	app := &Application{
		// adRepository: to be added
		dbAdapter: *dbAdapter,
	}

	return app
}

func (app *Application) Run() error {
	// Start the application logic
	// ...
	return nil
}
