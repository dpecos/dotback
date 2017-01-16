package utils

import (
	"github.com/fatih/color"
)

func Debug(msg string, args ...interface{}) {
	color.Blue(msg, args...)
}

func Info(msg string, args ...interface{}) {
	color.White(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	color.Magenta(msg, args...)
}

func Error(msg string, args ...interface{}) {
	color.Red(msg, args...)
}
