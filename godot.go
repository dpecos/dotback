package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dpecos/godot/models"
)

func ReadConfig() ([]models.Recipe, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("Could not read config.json file: %s", err)
	}

	var config []models.Recipe
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("Could not parse config.json file: %s", err)
	}

	return config, nil
}

func ExecRecipes(config []models.Recipe) {
	for _, recipe := range config {
		recipe.Exec()
	}
}

func main() {
	fmt.Printf("godot\n-----\n")

	config, err := ReadConfig()
	if err != nil {
		fmt.Println("Error loading config.json file\n", err)
		os.Exit(1)
	}

	ExecRecipes(config)
}
