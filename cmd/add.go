/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new place",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Take the values of place by flags
		place := db.Place{
			Name: args[0],
			Type: args[1],
			Last: time.Now(),
		}

		err := db.AddPlace(place)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("added to db")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
