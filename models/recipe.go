package models

import "fmt"

type Recipe struct {
	Name    string   `json:"name"`
	Actions []Action `json:"actions"`
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
