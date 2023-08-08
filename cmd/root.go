package cmd

import (
	"fmt"
	"os"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mekan",
	Short: "A brief description of your application",

	Run: func(cmd *cobra.Command, args []string) {
		if list, _ := cmd.Flags().GetBool("list"); list {
			printPlacesTable(db.AllPlaces())
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("list", "l", false, "List all places")
}

func printPlacesTable(places []db.Place) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Name", "Type", "Last"})

	for _, place := range places {

		tw.AppendRow(table.Row{
			place.Id, place.Name, place.Type, place.Last.Format("02/01/2006"),
		})
	}

	fmt.Println(tw.Render())
}
