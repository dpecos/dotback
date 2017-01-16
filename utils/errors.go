package utils

import "os"

func CheckError(msg string, err error) {
	if err == nil {
		return
	}
	if err != nil {
		Error("%s: %s", msg, err)
		os.Exit(-1)
	}
}
