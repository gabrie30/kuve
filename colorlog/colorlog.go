// Package colorlog has various Print functions that can be called to change the color of the text in standard out
package colorlog

import (
	"os"

	"github.com/fatih/color"
)

// Info prints yellow colored text to standard out
func Info(msg interface{}) {
	color.New(color.FgYellow).Println(msg)
}

// Success prints green colored text to standard out
func Success(msg interface{}) {
	color.New(color.FgGreen).Println(msg)
}

// Error prints red colored text to standard out
func Error(msg interface{}) {
	color.New(color.FgRed).Println(msg)
}

// FatalError prints red colored text to standard out and exits
func FatalError(msg interface{}) {
	color.New(color.FgRed).Println(msg)
	os.Exit(1)
}

// SubtleInfo prints magenta colored text to standard out
func SubtleInfo(msg interface{}) {
	color.New(color.FgHiMagenta).Println(msg)
}
