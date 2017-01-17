package steps

import (
	"fmt"

	. "github.com/dpecos/dotback/utils"
)

// Fetches and install go packages
func AptInstall(recipe string, num int, packages []string) error {
	if len(packages) == 0 {
		return nil
	}
	for _, pkg := range packages {
		fmt.Printf(" · [#%d apt-install] Installing apt package '%s'\n", num, pkg)
		err := Execute(fmt.Sprintf("sudo apt install --assume-yes %s", pkg))
		if err != nil {
			return err
		}
	}
	return nil
}
