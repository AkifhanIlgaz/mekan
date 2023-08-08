/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete place",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Please enter valid id")
				return
			}
			err = db.DeletePlace(id)
			if err != nil {
				fmt.Println("Cannot delete place: ", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
