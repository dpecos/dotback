package actions

import (
	"fmt"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type Cmd struct {
	models.Action
}

// Executes a shell command
func (cmd Cmd) Execute(recipe models.Recipe, pos int) error {
	command := cmd.Arguments[0]

	fmt.Printf(" · [#%d cmd] Executing %s\n", pos, command)
	return utils.Execute(command)
}
