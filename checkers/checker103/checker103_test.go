package checker103

import (
	. "github.com/tianshi25/additionalCheck/db"
	"reflect"
	"testing"
)

func TestCheckErrorCopyRightDateNotUpdated(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{"-2021", []string(nil)},
		{"/* 1993-2021 */", []string(nil)},
		{"/* 1993-2020 */", []string{"ERROR   file1:0:Please check copy right last modify date"}},
		{"-2020", []string(nil)},
		{"", []string(nil)},
	}

	for _, test := range tests {
		records := check("file1", test.in)
		out := GetRecordsStr(records)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test fail, input:%#v expected:%#v output:%#v",
				test.in, test.out, out)
		}
	}
}
