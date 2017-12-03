package utils

import "os"
import "fmt"

func CheckError(msg string, err error) {
	if err == nil {
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", msg, err)
		os.Exit(-1)
	}
}
