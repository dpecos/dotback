package steps

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/dpecos/godot/utils"
)

// Link creates links for files matching pattern in the recipe folder
func Link(recipe string, pattern string) error {
	from := path.Join(utils.HomeDir(), ".dotfiles", recipe)
	var err error
	if _, err = os.Stat(from); err == nil {
		if pattern == "*" {
			err = filepath.Walk(from, func(path string, file os.FileInfo, err error) error {
				if !file.IsDir() {
					to := fmt.Sprintf("/tmp/home/.%s", file.Name())
					fmt.Printf(" · Linking %s -> %s\n", path, to)
					return os.Symlink(path, to)
				}
				return nil
			})
		} else if pattern == "" {
			to := fmt.Sprintf("/tmp/home/.%s", recipe)
			fmt.Printf(" · Linking %s -> %s\n", from, to)
			err = os.Symlink(from, to)
		}
	}

	// fmt.Println(err)

	return err
}
