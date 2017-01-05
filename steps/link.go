package steps

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/dpecos/godot/utils"
)

func linkUndo(recipe string, num int, path string) {
	fmt.Printf(" · [#%d link] Unlinking %s\n", num, path)
	os.Remove(path)
}

// Link creates links for files matching pattern in the recipe folder
func Link(recipe string, num int, pattern string) error {
	if pattern == "" {
		return nil
	}
	from := path.Join(utils.HomeDir(), ".dotfiles", recipe)
	var err error
	if _, err = os.Stat(from); err == nil {
		if pattern == "*" {
			err = filepath.Walk(from, func(path string, file os.FileInfo, err error) error {
				if !file.IsDir() {
					to := fmt.Sprintf("/tmp/home/.%s", file.Name())
					linkUndo(recipe, num, to)
					fmt.Printf(" · [#%d link] Linking %s -> %s\n", num, path, to)
					return os.Symlink(path, to)
				}
				return nil
			})
		} else if pattern == "." {
			to := fmt.Sprintf("/tmp/home/.%s", recipe)
			linkUndo(recipe, num, to)
			fmt.Printf(" · [#%d link] Linking %s -> %s\n", num, from, to)
			err = os.Symlink(from, to)
		}
	}

	// fmt.Println(err)

	return err
}
