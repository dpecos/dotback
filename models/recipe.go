package models

import "fmt"

type Recipe struct {
	Name       string
	Attributes []string
	Actions    []Action
}

func (recipe Recipe) Exec() {
	fmt.Printf("\nExecuting recipe '%s'\n", recipe.Name)
	for i, action := range recipe.Actions {
		err := action.Exec(recipe.Name, i)
		if err != nil {
			fmt.Printf("   ERROR executing action #%d of recipe '%s': %s\n", i, recipe.Name, err)
			return
		}
	}
	return
}

func (recipe *Recipe) addAction(action Action) {
	recipe.Actions = append(recipe.Actions, action)
}

func (recipe *Recipe) ActionNames() []string {
	var actions []string
	for _, a := range recipe.Actions {
		actions = append(actions, a.Name)
	}
	return actions
}
