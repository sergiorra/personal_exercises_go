package ontario

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/errors"
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/models"
)

const (
	productsEndpoint = "/products"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type beerRepo struct {
	url string
}

// NewOntarioRepository fetch beers data from csv
func NewOntarioRepository() models.BeerRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []models.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEndpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", productsEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", productsEndpoint)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}
