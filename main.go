/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"path/filepath"

	"github.com/AkifhanIlgaz/mekan/cmd"
	"github.com/AkifhanIlgaz/mekan/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "places.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
