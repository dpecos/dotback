package models

import (
	"fmt"
	"reflect"

	"github.com/dpecos/godot/steps"
)

type Action struct {
	Link string `json:"link"`
	Cmd  string `json:"cmd"`
}

func (action Action) Exec(recipeName string) error {

	v := reflect.ValueOf(&action).Elem()
	typeAction := v.Type()

	for i := 0; i < v.NumField(); i++ {
		step := typeAction.Field(i).Name
		value := v.Field(i)

		var err error

		if value.Interface() != nil {
			switch step {
			case "Link":
				err = steps.Link(recipeName, value.String())
			case "Cmd":
			default:
				err = fmt.Errorf("Unknown action %+v", step)
			}

			if err != nil {

				return err
			}
		}
	}
	return nil
}
