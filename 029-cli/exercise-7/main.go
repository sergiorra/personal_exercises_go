
package main

import (
	"flag"

	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/beers"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/fetching"
	models "github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/models"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/repository/csv"
	"github.com/sergiorra/personal_exercises_go/029-cli/exercise-7/repository/ontario"
	"github.com/spf13/cobra"
)

func main() {
	// CPU profiling code starts here
	//f, _ := os.Create("beers.cpu.prof")
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	// CPU profiling code ends here

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

	// Memory profiling code starts here
	//f2, _ := os.Create("beers.mem.prof")
	//defer f2.Close()
	//pprof.WriteHeapProfile(f2)
	// Memory profiling code ends here
}