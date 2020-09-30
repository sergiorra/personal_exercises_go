package csv

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-3/models"
)

type repository struct {
}

// NewRepository initialize csv repository
func NewRepository() models.BeerRepo {
	return &repository{}
}

// GetBeers fetch beers data from csv
func (r *repository) GetBeers() ([]models.Beer, error) {
	f, _ := os.Open("data/beers.csv")
	reader := bufio.NewReader(f)

	var beers []models.Beer

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		productID, _ := strconv.Atoi(values[0])
		price, _ := strconv.ParseFloat(values[3], 64)

		beer := models.NewBeer(
			productID,
			values[1],
			values[2],
			values[5],
			values[6],
			models.NewBeerType(values[4]),
			price,
		)

		beers = append(beers, beer)
	}

	return beers, nil
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}