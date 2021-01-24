package checker003

import (
	"reflect"
	"testing"
	. "tianshi25.github.com/additionalCheck/db"
)

func TestCheckExtraSpace(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{"   ", []string{"ERROR   file1:1:extra space is detected at position 2"}},
		{"   \n", []string{"ERROR   file1:1:extra space is detected at position 2"}},
		{"    abc", []string(nil)},
		{"    abc ", []string{"ERROR   file1:1:extra space is detected at position 7"}},
		{"    a  bc", []string{"ERROR   file1:1:extra space is detected at position 5"}},
		{"    abc  //efg", []string(nil)},
		{"    abc  /*efg*/", []string(nil)},
		{"    abc  /*efg*/ ", []string{"ERROR   file1:1:extra space is detected at position 16"}},
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
