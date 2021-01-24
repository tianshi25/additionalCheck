package checker201

import (
	. "github.com/tianshi25/additionalCheck/db"
	"reflect"
	"testing"
)

func TestCheckErrorUsingTab(t *testing.T) {
	tests := []struct {
		fileName string
		in       string
		out      []string
	}{
		{"1.c", "line", []string(nil)},
		{"1.c", "line ||", []string(nil)},
		{"1.c", "line &&", []string(nil)},
		{"1.c", "  || line", []string{"ERROR   1.c:1:Operator should be located at line end"}},
		{"1.c", "  && line", []string{"ERROR   1.c:1:Operator should be located at line end"}},
		{"1.java", "line", []string(nil)},
		{"1.java", "line ||", []string{"ERROR   1.java:1:Operator should be located at line start"}},
		{"1.java", "line &&", []string{"ERROR   1.java:1:Operator should be located at line start"}},
		{"1.java", "  || line", []string(nil)},
		{"1.java", "  && line", []string(nil)},
		{"1.c", "", []string(nil)},
	}

	for _, test := range tests {
		records := check(test.fileName, test.in)
		out := GetRecordsStr(records)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test fail, filename %#v input:%#v expected:%#v output:%#v",
				test.fileName, test.in, test.out, out)
		}
	}
}
