package main

import (
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-3/beers"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-3/repository"
	"github.com/spf13/cobra"
)

func main() {
	// modeling beer data
	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd(csvRepo))
	rootCmd.Execute()
}