package models

import (
	"fmt"
)

type Command interface {
	Execute(recipe Recipe, pos int) error
	GetName() string
	GetArguments() []string
}

type Action struct {
	Name      string
	Arguments []string
}

func (action Action) Execute(recipe Recipe, pos int) error {
	return fmt.Errorf("Not implmented")
}

func (action Action) GetName() string {
	return action.Name
}

func (action Action) GetArguments() []string {
	return action.Arguments
}
