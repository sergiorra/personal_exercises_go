package main

import (
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-1/beers"
	"github.com/spf13/cobra"
)

func main() {
	// beers in program memory
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd())
	rootCmd.Execute()
}