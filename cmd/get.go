package cmd

import (
	"fmt"
	"github.com/TheLazyLemur/pokemon-cli/pokemon"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Pokemon by ID",
	Long: `Get a Pokemon by its ID`,
	Run: func(cmd *cobra.Command, args []string) {
		intVar, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		p, err := pokemon.GetById(intVar)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pokemon id: %+v\n", p.Id)
		fmt.Printf("Pokemon name: %+v\n", p.Species.Name)
		fmt.Printf("Pokemon weight: %+v\n", p.Weight)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
