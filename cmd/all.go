/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/spf13/cobra"
)

// TODO: Delete this command. Use -ls flag to list all places
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		allPlaces := db.AllPlaces()

		for _, place := range allPlaces {
			fmt.Println(place)
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
