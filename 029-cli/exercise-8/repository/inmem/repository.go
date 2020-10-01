package inmem

import (
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-8/models"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() models.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from memory
func (r *repository) GetBeers() ([]models.Beer, error) {
	return []models.Beer{
		models.NewBeer(
			127,
			"Mad Jack Mixer",
			"Domestic Specialty",
			"Molson",
			"Canada",
			"23.95",
			models.NewBeerType("Lager"),
		),
		models.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Grolsch Export B.V.",
			"Canada",
			"49.50",
			models.NewBeerType("Non-Alcoholic Beer"),
		),
	}, nil
}
