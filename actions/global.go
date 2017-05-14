package actions

import (
	"fmt"

	"github.com/dpecos/dotback/models"
)

func CreateAction(name string, arguments []string) (models.Command, error) {
	var action models.Command
	for i, arg := range arguments {
		if arg[0] == '"' {
			arguments[i] = arg[1 : len(arg)-1]
		}
	}
	switch name {
	case "link":
		link := Link{}
		link.Name = name
		link.Arguments = arguments
		action = link
	case "cmd":
		cmd := Cmd{}
		cmd.Name = name
		cmd.Arguments = arguments
		action = cmd
	case "apt":
		apt := Apt{}
		apt.Name = name
		apt.Arguments = arguments
		action = apt
	case "git":
		git := Git{}
		git.Name = name
		git.Arguments = arguments
		action = git
	case "go-get":
		goget := GoGet{}
		goget.Name = name
		goget.Arguments = arguments
		action = goget
	case "include":
		include := Include{}
		include.Name = name
		include.Arguments = arguments
		action = include
	default:
		err := fmt.Errorf("Command %s not implemented", name)
		return nil, err
	}

	return action, nil
}
