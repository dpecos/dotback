package tests

import (
	"testing"

	"github.com/dpecos/dotback/dotback"

	"reflect"
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

	if !reflect.DeepEqual(recipes, expected) {
		t.Errorf("Recipe list not matching expected (got %q, expected %q)", recipes, expected)
	}

}
func TestRetrieveActionsInRecipe(t *testing.T) {
	c := loadConfig(t)

	recipe, err := c.GetRecipe("bash")
	if err != nil {
		t.Fatalf("Recipe not found")
	}

	expected := []string{"link", "include"}
	actions := recipe.ActionNames()

	if !reflect.DeepEqual(actions, expected) {
		t.Errorf("Action list not matching expected (got %q, expected %q)", actions, expected)
	}

	expected = []string{"~/.bashrc", "source ~/.bash_include"}
	includeAction := recipe.Actions[1]
	arguments := includeAction.GetArguments()

	if !reflect.DeepEqual(arguments, expected) {
		t.Errorf("Argument list not matching expected (got %q, expected %q)", arguments, expected)
	}

}
func TestParametersInRecipe(t *testing.T) {
	c := loadConfig(t)

	recipe, err := c.GetRecipe("conky")
	if err != nil {
		t.Fatalf("Recipe not found")
	}

	expected := []string{"host=!nayar", "disabled"}
	attrs := recipe.Attributes

	if !reflect.DeepEqual(attrs, expected) {
		t.Errorf("Attribution list not matching expected (got %q, expected %q)", attrs, expected)
	}
}
