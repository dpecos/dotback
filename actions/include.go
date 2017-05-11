package actions

import (
	"github.com/dpecos/dotback/models"
)

type Include struct {
	models.Action
}

// Fetches and install go packages
func (include Include) Execute(recipe models.Recipe, pos int) error {
	return nil
}
