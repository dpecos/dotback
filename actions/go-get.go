package actions

import (
	"fmt"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type GoGet struct {
	models.Action
}

// Fetches and install go packages
func (goget GoGet) Execute(recipe models.Recipe, pos int) error {
	pkg := goget.Arguments[0]

	fmt.Printf(" · [#%d go-get] Installing go package '%s'\n", pos, pkg)
	err := utils.Execute(fmt.Sprintf("go get -u -v %s", pkg))
	if err != nil {
		return err
	}

	return nil
}
