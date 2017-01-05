package steps

import (
	"fmt"

	"path"

	"github.com/dpecos/dotback/utils"
)

// Clones a git repository
func Git(recipe string, num int, repo string) error {
	if repo == "" {
		return nil
	}
	fmt.Printf(" · [#%d git] Clonning git repo '%s'\n", num, repo)
	to := path.Join(utils.HomeDir(), "."+recipe)
	return utils.Execute(fmt.Sprintf("(rm -rf %s || true) && git clone %s %s", to, repo, to))
}
