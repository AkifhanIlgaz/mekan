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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
