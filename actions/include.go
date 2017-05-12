package actions

import (
	"log"

	"github.com/dpecos/dotback/models"
)

type Include struct {
	models.Action
}

// Includes the content of a file inside another
func (include Include) Execute(recipe models.Recipe, pos int) error {
	log.Fatalln("Not implemented")
	return nil
}
