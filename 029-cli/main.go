package main

import (
	"github.com/sergiorra/personal_exercises_go/029-cli/beers"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beers.InitBeersCmd())
	rootCmd.Execute()
}