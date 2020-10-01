package fetching

import (
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/models"
	"github.com/pkg/errors"
)

// Service provides beer fetching operations
type Service interface {
	// FetchBeers fetch all beers from repository
	FetchBeers() ([]models.Beer, error)
	// FetchByID filter all beers and get only the beer that match with given id
	FetchByID(id int) (models.Beer, error)
}

type service struct {
	bR models.BeerRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r models.BeerRepo) Service {
	return &service{r}
}

func (s *service) FetchBeers() ([]models.Beer, error) {
	return s.bR.GetBeers()
}

func (s *service) FetchByID(id int) (models.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return models.Beer{}, err
	}

	for _, beer := range beers {
		if beer.ProductID == id {
			return beer, nil
		}
	}

	return models.Beer{}, errors.Errorf("Beer %d not found", id)
}