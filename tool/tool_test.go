package tool

import (
    "testing"
    "reflect"
)

func TestRemoveWindowsLineEnd(t *testing.T) {
    tests := [] struct {
        in string
        out string
    } {
        {"abc\r\nefg", "abc\nefg"},
        {"abc\r\n\r\nefg", "abc\n\nefg"},
        {"abc\nefg", "abc\nefg"},
        {"abcefg", "abcefg"},
        {"", ""},
    }

    for _, test := range(tests) {
        out := RemoveWindowsLineEnd(test.in)
        if out != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, out)
        }
    }
}

func TestGetAllComments(t *testing.T) {
    tests := [] struct {
        in string
        out []string
    } {
        { "adb//comment data\n\nefg", []string{"//comment data" }, },
        { "adb/*comment\n *comment2*/efg", []string{ "/*comment\n *comment2*/" }, },
        { "adb//comment data\n\nefg/*comment\n *comment2*/", []string{"//comment data", "/*comment\n *comment2*/"}, },
        { "adbefg", []string{}, },
        { "/*comment\n *comment2*/adb//comment data\nef", []string{"/*comment\n *comment2*/", "//comment data"}, },
        { "", []string{}, },
    }

    for _, test := range(tests) {
        out := GetAllComments(test.in)
        if !reflect.DeepEqual(out, test.out) {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, out)
        }
    }
}

func TestGetReplacementForComments(t *testing.T) {
    tests := [] struct {
        in string
        out string
    } {
        { "adb\n\nefg", "\n\n" },
        { "adb\n\nef\ng", "\n\n\n" },
        { "adbefg", "" },
        { "", "" },
    }

    for _, test := range(tests) {
        ret := string(getReplacementForComments([]byte(test.in)))
        if ret != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, ret)
        }
    }
}

func TestReplaceComments(t *testing.T) {
    tests := [] struct {
        in string
        out string
    } {
        { "adb\n   /*a\ndb*/efg\nhi", "adb\n\nefg\nhi" },
        { "adb    //comment\n\nef\ng", "adb\n\nef\ng" },
        { "adbefg", "adbefg" },
        { "", "" },
    }

    for _, test := range(tests) {
        if ReplaceComments(test.in) != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, ReplaceComments(test.in))
        }
    }
}