package fetching

import (
	"errors"
	"testing"

	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-9/models"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-9/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestFetchByID(t *testing.T) {

	tests := map[string]struct {
		repo  models.BeerRepo
		input int
		want  int
		err   error
	}{
		"valid beer":            {repo: buildMockBeers(), input: 127, want: 127, err: nil},
		"not found beer":        {repo: buildMockBeers(), input: 99999, err: errors.New("error")},
		"error with repository": {repo: buildMockError(), err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := NewService(tc.repo)

			b, err := service.FetchByID(tc.input)

			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.ProductID)
		})
	}
}

func buildMockBeers() models.BeerRepo {
	mockedRepo := &mock.BeerRepoMock{
		GetBeersFunc: func() ([]models.Beer, error) {
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
		},
	}

	return mockedRepo
}

func buildMockError() models.BeerRepo {
	mockedRepo := &mock.BeerRepoMock{
		GetBeersFunc: func() ([]models.Beer, error) {
			return []models.Beer{}, errors.New("error")
		},
	}

	return mockedRepo
}
