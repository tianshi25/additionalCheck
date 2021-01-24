package para

import (
	"reflect"
	"regexp"
	"testing"
)

func TestGetIgnoreCheckerId(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{
			"1,2,3",
			[]int{1, 2, 3},
		},
		{
			"1",
			[]int{1},
		},
		{
			"",
			[]int{},
		},
		{
			"1,  2,3,adb,,",
			[]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		ret := getIgnoreCheckerId(test.in)
		if reflect.DeepEqual(ret, tests) {
			t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
		}
	}
}

func TestGetExtensionList(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			"c,h,cpp,hpp,java,go",
			[]string{"c", "h", "cpp", "hpp", "java", "go"},
		},
		{
			"",
			[]string{},
		},
	}

	for _, test := range tests {
		ret := getExtensionList(test.in)
		if reflect.DeepEqual(ret, tests) {
			t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
		}
	}
}

func TestPathRegex(t *testing.T) {
	tests := []struct {
		extensions []string
		path       string
		result     bool
	}{
		{
			[]string{"c", "h", "cpp", "hpp", "java", "go"},
			"1.c",
			true,
		},
		{
			[]string{"c", "h", "cpp", "hpp", "java", "go"},
			"1.c1",
			false,
		},
		{
			[]string{"c", "h", "cpp", "hpp", "java", "go"},
			"1.",
			false,
		},
	}

	for _, test := range tests {
		regex := getExtensionRegex(test.extensions)
		matched, err := regexp.MatchString(regex, test.path)
		if matched != test.result || err != nil {
			t.Errorf("test fail, test:%#v | regex:%v matched:%#v error:%#v", test, regex, matched, err)
		}
	}
}
