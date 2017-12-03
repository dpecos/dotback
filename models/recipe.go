package models

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/dpecos/dotback/utils"
)

type Recipe struct {
	Name       string
	Attributes []string
	Actions    []Command
}

func (recipe Recipe) Execute() {
	if ex, reason := recipe.shouldExecute(); !ex {
		fmt.Printf("\nRecipe '%s' not being executed: %s\n", recipe.Name, reason)
		return
	}
	fmt.Printf("\nExecuting recipe '%s'\n", recipe.Name)
	for i, action := range recipe.Actions {
		err := action.Execute(recipe, i+1)
		if err != nil {
			fmt.Printf("ERROR executing action #%d of recipe '%s': %s\n", i, recipe.Name, err)
			return
		}
	}
	return
}

func (recipe *Recipe) shouldExecute() (bool, string) {
	for _, attr := range recipe.Attributes {
		if attr == "disabled" {
			return false, "Recipe disabled"
		}

		host, err := os.Hostname()
		utils.CheckError("Could not retrieve host name", err)

		if ok, reason := check(attr, "host", host); !ok {
			return false, reason
		}

		goos := runtime.GOOS
		if ok, reason := check(attr, "os", goos); !ok {
			return false, reason
		}
	}

	return true, ""
}

func check(attr, param, value string) (bool, string) {
	if strings.HasPrefix(attr, param+"=") {
		if strings.HasSuffix(attr, "=!"+value) {
			return false, param + " excluded (" + value + "): " + attr
		}
		if !strings.Contains(attr, "=!") && !strings.HasSuffix(attr, "="+value) {
			return false, param + " does not match (" + value + "): " + attr
		}
	}

	return true, ""
}

func (recipe *Recipe) AddAction(action Command) {
	recipe.Actions = append(recipe.Actions, action)
}

func (recipe *Recipe) ActionNames() []string {
	var actions []string
	for _, a := range recipe.Actions {
		if a != nil && a.GetName() != "" {
			actions = append(actions, a.GetName())
		}
	}
	return actions
}
