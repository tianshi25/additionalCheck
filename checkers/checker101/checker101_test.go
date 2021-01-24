package checker101

import (
	"reflect"
	"testing"
	. "tianshi25.github.com/additionalCheck/db"
)

func TestCheckMixCommentType(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{"//line1\n//line2", []string(nil)},
		{"/*comm\nen\nt1*///line2", []string{"NOTE    file1:0:mixing two comment type"}},
		{"//line2\n/*comm\nen\nt1*/", []string{"NOTE    file1:0:mixing two comment type"}},
		{"/*comm\nen\nt1*//*comm\nen\nt2*/", []string(nil)},
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
