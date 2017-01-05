package utils

import (
	"fmt"
	"os/exec"
)

func Execute(command string) error {
	out, err := exec.Command("sh", "-c", command).Output()
	fmt.Printf("%s", out)

	return err
}
