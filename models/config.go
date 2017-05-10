package models

import (
	"bufio"
	"log"
	"os"
	"strings"

	"fmt"
	"regexp"
)

type Config struct {
	Recipes []Recipe
}

func matches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)
	if len(match) != 0 {
		result := make(map[string]string)
		for i, name := range regex.SubexpNames() {
			if i != 0 {
				result[name] = match[i]
			}
		}
		return result
	}
	return nil
}

func (c *Config) Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	recipeRegex := regexp.MustCompile(`^\[(?P<name>[\w-]+)(\s+(?P<attrs>.*))?\]`)
	actionRegex := regexp.MustCompile(`^(?P<name>[\w-]+)(\s+(?P<args>.*))?`)
	argsRegex := regexp.MustCompile(`"[^"]*"|[^\s]+`)

	scanner := bufio.NewScanner(f)
	var recipe *Recipe
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) != 0 && line[0] != '#' {
			result := matches(recipeRegex, line)
			if result != nil {
				if recipe != nil {
					c.addRecipe(*recipe)
				}
				attrs := strings.Split(result["attrs"], " ")
				recipe = &Recipe{Name: result["name"], Attributes: attrs}
			} else {
				result = matches(actionRegex, line)
				if result != nil {
					args := argsRegex.FindAllString(result["args"], -1)
					action := Action{Name: result["name"], Arguments: args}
					recipe.addAction(action)
				}
			}
		}
	}
	// add last recipe to the config
	if recipe != nil {
		c.addRecipe(*recipe)
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

func (c Config) RecipeNames() []string {
	var names []string
	for _, r := range c.Recipes {
		names = append(names, r.Name)
	}
	return names
}

func (c Config) GetRecipe(name string) (*Recipe, error) {
	for _, r := range c.Recipes {
		if r.Name == name {
			return &r, nil
		}
	}

	return nil, fmt.Errorf("Recipe %s not found", name)
}
