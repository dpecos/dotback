package actions

import (
	"fmt"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type Apt struct {
	models.Action
}

// Fetches and install Debian / Ubuntu packages
func (apt Apt) Execute(recipe models.Recipe, pos int) error {
	packages := apt.Arguments[0]

	for _, pkg := range packages {
		fmt.Printf(" · [#%d apt-install] Installing apt package '%s'\n", pos, pkg)
		err := utils.Execute(fmt.Sprintf("sudo apt install --assume-yes %s", pkg))
		if err != nil {
			return err
		}
	}
	return nil
}
