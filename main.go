/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/AkifhanIlgaz/mekan/cmd"
	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "places.db")
	placeService, err := db.Init(dbPath, 24*time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println(placeService.TimePeriod)

	cmd.Execute()
}
