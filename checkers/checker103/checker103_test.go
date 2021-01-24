package checker103

import (
    . "tianshi25.github.com/additionalCheck/db"
    "testing"
    "reflect"
)

func TestCheckErrorCopyRightDateNotUpdated(t *testing.T) {
    tests := [] struct {
        in string
        out []string
    } {
        { "-2021", []string(nil) },
        { "/* 1993-2021 */", []string(nil) },
        { "/* 1993-2020 */", []string{"ERROR   file1:0:Please check copy right last modify date"} },
        { "-2020", []string(nil) },
        { "", []string(nil) },
    }

    for _, test := range(tests) {
        records := check("file1", test.in)
        out := GetRecordsStr(records)
        if !reflect.DeepEqual(out, test.out) {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, out)
        }
    }
}