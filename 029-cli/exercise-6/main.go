package main

import (
	"flag"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/fetching"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/repository/ontario"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/beers"
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/models"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-6/repository/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()
	var repo models.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd(fetchingService))
	rootCmd.Execute()
}
