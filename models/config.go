package models

import (
	"bufio"
	"log"
	"os"

	"regexp"
)

type Config struct {
	Recipes []Recipe
}

func (c *Config) Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	recipeRegex := regexp.MustCompile(`^\[(?P<name>[^\s]+).*\]`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		match := recipeRegex.FindStringSubmatch(line)
		if len(match) != 0 {
			result := make(map[string]string)
			for i, name := range recipeRegex.SubexpNames() {
				if i != 0 {
					result[name] = match[i]
				}
			}
			recipe := Recipe{Name: result["name"]}
			c.addRecipe(recipe)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *Config) addRecipe(r Recipe) {
	c.Recipes = append(c.Recipes, r)
}

func (c *Config) RecipeNames() []string {
	var names []string
	for _, r := range c.Recipes {
		names = append(names, r.Name)
	}
	return names
}
