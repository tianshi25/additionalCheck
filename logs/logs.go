package logs

import (
    "github.com/gookit/color"
    "fmt"
)

const (
    Verbose = 0
    Info = 1
    Warn = 2
    Error = 3
)

var level = Info

func SetLevel(newLevel int) {
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

func RecordSprintf(t int, format string, s ...interface{}) string {
    if t == Error {
        return fmt.Sprintf("Error: " + format, s...)
    }
    if t == Warn {
        return fmt.Sprintf("Warn:  " + format, s...)
    }
    return fmt.Sprintf("Info:  " + format, s...)
}