package db

import (
	"reflect"
	"testing"
)

func TestLess(t *testing.T) {
	tests := []struct {
		r1  Record
		r2  Record
		out bool
	}{
		{Record{"path1", 1, 1, []interface{}{}}, Record{string("path2"), 1, 1, []interface{}{}}, true},
		{Record{"path2", 1, 1, []interface{}{}}, Record{string("path1"), 1, 1, []interface{}{}}, false},
		{Record{"path1", 1, 1, []interface{}{}}, Record{string("path1"), 2, 1, []interface{}{}}, true},
		{Record{"path1", 2, 1, []interface{}{}}, Record{string("path1"), 1, 1, []interface{}{}}, false},
		{Record{"path1", 1, 1, []interface{}{}}, Record{string("path1"), 1, 2, []interface{}{}}, true},
		{Record{"path1", 1, 2, []interface{}{}}, Record{string("path1"), 1, 1, []interface{}{}}, false},
		{Record{"path1", 1, 1, []interface{}{}}, Record{string("path1"), 1, 1, []interface{}{"para1"}}, true},
		{Record{"path1", 1, 1, []interface{}{"para1"}}, Record{string("path1"), 1, 1, []interface{}{}}, false},
	}

	for _, test := range tests {
		ret := test.r1.Less(test.r2)
		if ret != test.out {
			t.Errorf("test fail, input:%#v %#v expected:%#v output:%#v",
				test.r1, test.r2, test.out, ret)
		}
	}
}

func TestSwap(t *testing.T) {
	r1 := Record{"path1", 1, 1, []interface{}{"para1"}}
	r1_copy := Record{"path1", 1, 1, []interface{}{"para1"}}
	r2 := Record{"path2", 2, 2, []interface{}{"para2"}}
	r2_copy := Record{"path2", 2, 2, []interface{}{"para2"}}

	r1.Swap(&r2)
	if !reflect.DeepEqual(r2, r1_copy) || !reflect.DeepEqual(r1, r2_copy) {
		t.Errorf("test fail")
	}
}
