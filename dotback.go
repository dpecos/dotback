package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/dpecos/dotback/models"
	. "github.com/dpecos/dotback/utils"
)

func ReadConfig(recipe string) []models.Recipe {
	file, err := ioutil.ReadFile(path.Join(HomeDir(), ".dotfiles", "config.json"))
	CheckError("Could not read config.json file", err)

	var config []models.Recipe
	err = json.Unmarshal(file, &config)
	CheckError("Could not parse config.json file", err)

	if recipe != "" {
		oldConfig := config
		config = nil

		for _, r := range oldConfig {
			if r.Name == recipe {
				fmt.Printf("Executing only recipe '%s' (skipping the rest)\n", r.Name)
				config = []models.Recipe{r}
			}
		}

		if config == nil {
			Error("Recipe %s not found", recipe)
			os.Exit(-1)
		}
	}

	return config
}

func ExecRecipes(config []models.Recipe) {
	for _, recipe := range config {
		recipe.Exec()
	}
}

var (
	app = kingpin.New("dotback", "Handle your dot files like a boss")
	// initialize = app.Command("init", "Use a Git repository to initialize your dotfiles folder")
	pull    = app.Command("pull", "Fetch latest changes from remote dotfiles repository")
	push    = app.Command("push", "Send latest changes to remote dotfiles repository")
	list    = app.Command("list", "List the actions defined in your ~/.dotfiles/config.json")
	install = app.Command("install", "Performs the actions defined in your ~/.dotfiles/config.json")
	recipe  = install.Arg("recipe", "Execute only this recipe").String()
	// add        = app.Command("add", "Creates a new recipe")
	// delete     = app.Command("delete", "Remove a recipe")
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "pull":
		path := path.Join(HomeDir(), ".dotfiles")
		err := Execute(fmt.Sprintf("cd %s && git pull", path))
		CheckError("Error updating repository", err)
	case "push":
		path := path.Join(HomeDir(), ".dotfiles")
		err := Execute(fmt.Sprintf("cd %s && git push", path))
		CheckError("Error updating repository", err)
	case "list":
		config := ReadConfig(*recipe)
		for _, recipe := range config {
			fmt.Printf("%s -> \n", recipe.Name)

			for _, action := range recipe.Actions {
				if action.Link != "" {
					fmt.Printf("   Link: %s\n", action.Link)
				}
				if action.Cmd != "" {
					fmt.Printf("   Cmd: %s\n", action.Cmd)
				}
				if action.Git != "" {
					fmt.Printf("   Git: %s\n", action.Git)
				}
			}
		}
	case "install":
		config := ReadConfig(*recipe)
		ExecRecipes(config)
	}

}
