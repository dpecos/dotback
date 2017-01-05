package steps

import "fmt"
import "github.com/dpecos/dotback/utils"

// Executes a shell command
func Cmd(recipe string, num int, command string) error {
	if command == "" {
		return nil
	}
	fmt.Printf(" · [#%d cmd] Executing '%s'\n", num, command)
	return utils.Execute(command)
}
