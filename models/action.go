package models

import (
	"fmt"
	"reflect"

	"github.com/dpecos/dotback/steps"
	"github.com/dpecos/dotback/utils"
)

type Action struct {
	Name      string
	Arguments []string
}

func (action Action) Exec(recipeName string, num int) error {

	v := reflect.ValueOf(&action).Elem()
	typeAction := v.Type()

	for i := 0; i < v.NumField(); i++ {
		step := typeAction.Field(i).Name
		value := v.Field(i)

		var err error

		if value.Interface() != nil {
			switch step {
			case "Link":
				err = steps.Link(recipeName, num, value.String())
			case "Cmd":
				err = steps.Cmd(recipeName, num, value.String())
			case "Git":
				err = steps.Git(recipeName, num, value.String())
			case "GoGet":
				err = steps.GoGet(recipeName, num, value.Interface().([]string))
			case "AptInstall":
				err = steps.AptInstall(recipeName, num, value.Interface().([]string))
			default:
				err = fmt.Errorf("Unknown action %+v", step)
			}

			utils.CheckError("Error executing action", err)
		}
	}
	return nil
}
