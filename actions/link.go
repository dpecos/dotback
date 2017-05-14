package actions

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type Link struct {
	models.Action
}

func linkUndo(num int, path string) {
	fmt.Printf(" · [#%d link] Unlinking %s\n", num, path)
	os.Remove(path)
}

// Link creates links for files matching pattern in the recipe folder
func (link Link) Execute(recipe models.Recipe, pos int) error {
	pattern := link.Arguments[0]

	from := path.Join(utils.HomeDir(), ".dotfiles", recipe.Name)
	var err error
	if _, err = os.Stat(from); err == nil {
		if pattern == "*" {
			err = filepath.Walk(from, func(filePath string, file os.FileInfo, err error) error {
				if !file.IsDir() {
					to := path.Join(utils.HomeDir(), fmt.Sprintf(".%s", file.Name()))
					linkUndo(pos, to)
					fmt.Printf(" · [#%d link] Linking %s -> %s\n", pos, filePath, to)
					return os.Symlink(filePath, to)
				}
				return nil
			})
		} else if pattern == "." {
			to := path.Join(utils.HomeDir(), fmt.Sprintf(".%s", recipe.Name))
			linkUndo(pos, to)
			fmt.Printf(" · [#%d link] Linking %s -> %s\n", pos, from, to)
			err = os.Symlink(from, to)
		}
	}

	return err
}
