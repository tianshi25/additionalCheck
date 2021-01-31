package logs

import (
	"github.com/gookit/color"
)

const (
	Verbose = 0
	Info    = 1
	Warn    = 2
	Error   = 3
	Invalid = 100
)

var level int

func init() {
	level = Warn
}

func SetLevel(newLevel int) {
	if newLevel == Invalid {
		E("Invalid log level")
	}
	level = newLevel
}

func V(format string, s ...interface{}) {
	format = format + "\n"
	if level <= Verbose {
		color.Note.Printf(format, s...)
	}
}

func I(format string, s ...interface{}) {
	format = format + "\n"
	if level <= Info {
		color.Info.Printf(format, s...)
	}
}

func W(format string, s ...interface{}) {
	format = format + "\n"
	if level <= Warn {
		color.Warn.Printf(format, s...)
	}
}

func E(format string, s ...interface{}) {
	format = format + "\n"
	if level <= Error {
		color.Error.Printf(format, s...)
	}
}
