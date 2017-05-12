package models

import "fmt"

type Recipe struct {
	Name       string
	Attributes []string
	Actions    []Command
}

func (recipe Recipe) Execute() {
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
