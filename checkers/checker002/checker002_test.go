package checker002

import (
	"reflect"
	"testing"
	. "tianshi25.github.com/additionalCheck/db"
)

func TestRuleIndentIncreaseNot4(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{"     abc\n", []string{"SUGGEST file1:1:indent should be multiples of 4 and increase by 4"}},
		{"     abc\n        efg\n", []string{"SUGGEST file1:1:indent should be multiples of 4 and increase by 4"}},
		{"    abc\n}efg\n", []string(nil)},
		{"    abc\n        efg\n", []string(nil)},
		{"    abc\n            efg\n", []string{"SUGGEST file1:2:indent should be multiples of 4 and increase by 4"}},
		{"    abc\n         efg\n", []string{"SUGGEST file1:2:indent should be multiples of 4 and increase by 4"}},
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
