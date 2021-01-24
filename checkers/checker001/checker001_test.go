package checker001

import (
	"reflect"
	"testing"
	. "tianshi25.github.com/additionalCheck/db"
)

func TestCheckTabDetected(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			"line1ab\t\nline2\nline3\t\n",
			[]string{
				"ERROR   file1:1:tab is detected at position 7",
				"ERROR   file1:3:tab is detected at position 5",
			},
		},
		{
			"\n\t\t",
			[]string{
				"ERROR   file1:2:tab is detected at position 0",
			},
		},
		{"line1\nline2\nline3\n", []string(nil)},
		{"", []string(nil)},
		{"\n\n\n", []string(nil)},
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
