package main

import (
	"flag"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-5/repository/ontario"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-5/beers"
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-5/models"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-5/repository/csv"
	"github.com/spf13/cobra"
)

func main() {
	// handling behavioural errors
	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()
	var repo models.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd(repo))
	rootCmd.Execute()
}
