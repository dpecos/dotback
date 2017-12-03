package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dpecos/dotback/models"
	"github.com/dpecos/dotback/utils"
)

type Include struct {
	models.Action
}

// Includes the content of a file inside another
func (include Include) Execute(recipe models.Recipe, pos int) error {
	f := utils.ResolveFile(include.Arguments[0])
	content := include.Arguments[1]

	fmt.Printf(" · [#%d include] Adding content to file '%s'\n", pos, f)

	found, err := include.checkContentInFile(f, content)
	if err != nil {
		return err
	}

	if found {
		fmt.Printf(" · [#%d include] Content already found in file. Skipping...\n", pos)
		return nil
	}

	err = include.addContentToFile(f, content)
	return err
}

func (include Include) checkContentInFile(fName string, content string) (bool, error) {
	f, err := ioutil.ReadFile(fName)
	if err != nil {
		return false, err
	}

	found := strings.Contains(string(f), content)
	return found, err
}

func (include Include) addContentToFile(fName string, data string) error {
	content, err := ioutil.ReadFile(fName)
	if err != nil {
		return err
	}

	content = append(content, []byte("\n"+data)...)

	err = ioutil.WriteFile(fName, content, os.ModeAppend)
	return err
}
