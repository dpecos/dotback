package tests

import (
	"testing"

	"github.com/dpecos/dotback/dotback"
)

func loadConfig(t *testing.T) (config dotback.Config) {
	if err := config.Load("config"); err != nil {
		t.Errorf("Error loading config: %s", err)
	}
	return
}

func TestRetrieveRecipes(t *testing.T) {
	c := loadConfig(t)

	expected := []string{"ssh", "git", "mercurial", "tmux", "bash", "conky", "nodejs", "vim", "go", "install-packages-deb"}
	recipes := c.RecipeNames()

	equals(t, expected, recipes)
}
func TestRetrieveActionsInRecipe(t *testing.T) {
	c := loadConfig(t)

	recipe, err := c.GetRecipe("bash")
	ok(t, err)

	expected := []string{"link", "include"}
	actions := recipe.ActionNames()

	equals(t, expected, actions)

	expected = []string{"~/.bashrc", "source ~/.bash_include"}
	includeAction := recipe.Actions[1]
	arguments := includeAction.GetArguments()

	equals(t, expected, arguments)
}
func TestParametersInRecipe(t *testing.T) {
	c := loadConfig(t)

	recipe, err := c.GetRecipe("conky")
	ok(t, err)

	expected := []string{"host=!nayar", "disabled"}
	attrs := recipe.Attributes

	equals(t, expected, attrs)
}
