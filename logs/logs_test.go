package logs

import "testing"
import "fmt"

func TestSetLevel(t *testing.T) {
    SetLevel(Verbose)
}

func TestE(t *testing.T) {
    E("Error message")
}

func TestW(t *testing.T) {
    W("Warning message")
}

func TestI(t *testing.T) {
    I("Info message")
}

func TestV(t *testing.T) {
    V("Verbose message")
}

func TestRecordSprintf(t *testing.T) {
    fmt.Println(RecordSprintf(Error, "error string"))
    fmt.Println(RecordSprintf(Warn, "warn string"))
    fmt.Println(RecordSprintf(Info, "info string"))
}
