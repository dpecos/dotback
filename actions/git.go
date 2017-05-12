package actions

import (
	"fmt"

	"path"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type Git struct {
	models.Action
}

// Clones a git repository
func (git Git) Execute(recipe models.Recipe, pos int) error {
	repo := git.Arguments[0]

	fmt.Printf(" · [#%d git] Clonning git repo %s\n", pos, repo)
	to := path.Join(utils.HomeDir(), "."+recipe.Name)
	return utils.Execute(fmt.Sprintf("(rm -rf %s || true) && git clone %s %s", to, repo, to))
}
