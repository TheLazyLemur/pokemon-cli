package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"github.com/TheLazyLemur/pokemon-cli/pokemon"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random Pokemon",
	Long: `Get a random Pokemon's stats and display it in the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().Unix())

		rangeLower := 1
		rangeUpper := pokemon.GetTotalPokemon()
		randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)

		p,err := pokemon.GetById(randomNum)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("Pokemon id: %+v\n", p.Id)
		fmt.Printf("Pokemon name: %+v\n", p.Species.Name)
		fmt.Printf("Pokemon weight: %+v\n", p.Weight)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
