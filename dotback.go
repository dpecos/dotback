package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

func ReadConfig() ([]models.Recipe, error) {
	file, err := ioutil.ReadFile(path.Join(utils.HomeDir(), ".dotfiles", "config.json"))
	if err != nil {
		return nil, fmt.Errorf("Could not read config.json file: %s", err)
	}

	var config []models.Recipe
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("Could not parse config.json file: %s", err)
	}

	// fmt.Printf("%+v", config)

	return config, nil
}

func ExecRecipes(config []models.Recipe) {
	for _, recipe := range config {
		recipe.Exec()
	}
}

func main() {
	fmt.Printf("dotback\n-----\n")

	config, err := ReadConfig()
	if err != nil {
		fmt.Println("Error loading config.json file\n", err)
		os.Exit(1)
	}

	ExecRecipes(config)
}
