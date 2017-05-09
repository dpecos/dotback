package tests

import (
	"testing"

	"github.com/dpecos/dotback/models"

	"reflect"
)

func TestRetrieveRecipes(t *testing.T) {
	expected := []string{"ssh", "git", "mercurial", "tmux", "bash", "conky", "nodejs", "vim", "go", "install-packages-deb"}

	var c models.Config
	if err := c.Load("config"); err != nil {
		t.Errorf("Error loading config: %s", err)
	}
	recipes := c.RecipeNames()

	if !reflect.DeepEqual(recipes, expected) {
		t.Errorf("Recipe list not matching expected (got %q, expected %q)", recipes, expected)
	}

}
