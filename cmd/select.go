package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select random place",
	Run: func(cmd *cobra.Command, args []string) {
		placeType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("select place: %w", err)
			return
		}

		places := filter(db.AllPlaces(), placeType)
		place, err := selectPlace(places)
		if err != nil {
			fmt.Println(err)
			return
		}

		printPlace(place)

		if confirm() {
			db.UpdateLast(place.Id)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().StringP("type", "t", "food", "filter by type")
}

func selectPlace(places []db.Place) (db.Place, error) {
	if len(places) == 0 {
		return db.Place{}, errors.New("there is no place to select. Please add place")
	}

	s := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(s)

	place := places[r.Intn(len(places))]

	return place, nil
}

func printPlace(place db.Place) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Name", "Type", "Last"})

	tw.AppendRow(table.Row{
		place.Id, strings.ToTitle(place.Name), place.Type, place.Last.Format("02/01/2006"),
	})

	fmt.Println(tw.Render())
}

func filter(places []db.Place, placeType string) []db.Place {
	var filteredPlaces []db.Place

	for _, place := range places {
		if place.Type == placeType {
			filteredPlaces = append(filteredPlaces, place)
		}
	}

	return filteredPlaces
}

func confirm() bool {
	fmt.Print("Do you want to go to this place (y/n) ? ")
	scanner.Scan()

	switch strings.ToLower(scanner.Text()) {
	case "y":
		return true
	default:
		return false
	}

}
