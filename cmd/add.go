package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new place",

	Run: func(cmd *cobra.Command, args []string) {
		place := getPlaceInfo()
		err := db.AddPlace(place)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

var scanner = bufio.NewScanner(os.Stdin)

func getPlaceInfo() db.Place {
	return db.Place{
		Name: getName(),
		Type: getType(),
		Last: time.Now(),
	}
}

func getName() string {
	fmt.Print("> Name: ")
	scanner.Scan()

	name := scanner.Text()
	if name == "" {
		fmt.Println("Please enter a name")
		return getName()
	}

	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	return name
}

func getType() string {
	fmt.Print("> Type: ")
	scanner.Scan()

	placeType := scanner.Text()
	if placeType == "" {
		fmt.Println("Please enter a placeType")
		return getType()
	}

	placeType = strings.TrimSpace(placeType)
	placeType = strings.ToLower(placeType)

	return placeType
}
