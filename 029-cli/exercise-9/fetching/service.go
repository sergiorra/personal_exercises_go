package fetching

import (
	"math"
	"sync"

	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-9/models"
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

	beersPerRoutine := 10
	numRoutines := numOfRoutines(len(beers), beersPerRoutine)

	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)

	var b models.Beer

	for i := 0; i < numRoutines; i++ {
		go func(id, begin, end int, beers []models.Beer, b *models.Beer, wg *sync.WaitGroup) {
			for i := begin; i <= end; i++ {
				if beers[i].ProductID == id {
					*b = beers[i]
				}
			}
			wg.Done()
		}(id, i, i+beersPerRoutine, beers, &b, wg)
	}

	wg.Wait()

	return b, nil
}

func numOfRoutines(numOfBeers, beersPerRoutine int) int {
	return int(math.Ceil(float64(numOfBeers) / float64(beersPerRoutine)))
}
