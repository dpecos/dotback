package dotback

import (
	"fmt"
	"path"

	"github.com/dpecos/dotback/utils"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

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

func readConfig() Config {
	file := path.Join(utils.HomeDir(), ".dotfiles", "config")

	var config Config
	config.Load(file)

	return config
}

func executeRecipes(config Config, recipeName *string) {
	for _, recipe := range config.Recipes {
		if recipeName != nil {
			if *recipeName == recipe.Name {
				recipe.Execute()
			}
		} else {
			recipe.Execute()
		}
	}
}

func Run(args []string) {
	switch kingpin.MustParse(app.Parse(args[1:])) {
	case "pull":
		path := path.Join(utils.HomeDir(), ".dotfiles")
		err := utils.Execute(fmt.Sprintf("cd %s && git pull", path))
		utils.CheckError("Error updating repository", err)
	case "push":
		path := path.Join(utils.HomeDir(), ".dotfiles")
		err := utils.Execute(fmt.Sprintf("cd %s && git push", path))
		utils.CheckError("Error updating repository", err)
	case "list":
		config := readConfig()
		for _, recipe := range config.Recipes {
			fmt.Printf("%s %s -> \n", recipe.Name, recipe.Attributes)

			for _, action := range recipe.Actions {
				fmt.Printf("   %s: %s\n", action.GetName(), action.GetArguments())
			}
			fmt.Println("")
		}
	case "install":
		config := readConfig()
		executeRecipes(config, recipe)
	}
}
