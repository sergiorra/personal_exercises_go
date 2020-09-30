package main

import (
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-2/beers"
	"github.com/spf13/cobra"
)

func main() {
	// extracting beers from CSV file
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd())
	rootCmd.Execute()
}