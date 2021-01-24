package checker102

import (
	. "github.com/tianshi25/additionalCheck/db"
	"reflect"
	"testing"
)

func TestCheckErrorMultilineCommitFormatWrong(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{"/*\n * comment1\n * comment2\n */", []string(nil)},
		{"    /*\n     * comment1\n     * comment2\n */", []string(nil)},
		{"    /* comment2 */", []string{"SUGGEST file1:1:type /*\\n * \\n */ multiline commit is suggested"}},
		{"    /* \ncomment2 */", []string{"SUGGEST file1:1:type /*\\n * \\n */ multiline commit is suggested"}},
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
