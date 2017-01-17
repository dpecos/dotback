package steps

import (
	"fmt"

	. "github.com/dpecos/dotback/utils"
)

// Fetches and install go packages
func GoGet(recipe string, num int, packages []string) error {
	if len(packages) == 0 {
		return nil
	}
	for _, pkg := range packages {
		fmt.Printf(" · [#%d go-get] Installing go package '%s'\n", num, pkg)
		err := Execute(fmt.Sprintf("go get -u -v %s", pkg))
		if err != nil {
			return err
		}
	}
	return nil
}
