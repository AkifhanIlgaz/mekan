/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select random place",
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := cmd.Flags().GetString("type")
		fmt.Println(s)

		fmt.Println("select called")
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().StringP("type", "t", "", "filter by type")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
