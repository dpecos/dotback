package utils

import (
	"log"
	"os/user"
	"strings"
)

func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func ResolveFile(fName string) string {
	if strings.Contains(fName, "~") {
		return strings.Replace(fName, "~", HomeDir(), -1)
	}
	return fName
}
