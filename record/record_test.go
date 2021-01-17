package record

import (
    "testing"
    "reflect"
)

func TestGetStr(t *testing.T) {
    tests := [] struct {
        r Record
        out string
    } {
        { Record{"path1", 1, ErrorUsingTab, []string{"11"}}, "Error: path1:1:tab is detected at position 11"},
        { Record{"path2", 2, ErrorWrongIndent, []string{"12"}}, "Warn:  path2:2:indent should be 12"},
        { Record{"path3", 3, ErrorExtraSpace, []string{"13"}}, "Warn:  path3:3:multiple space is detected at position 13"},
        { Record{"path4", 4, ErrorMixComment, []string{}}, "Warn:  path4:4:mixing two comment type"},
        { Record{"path5", 5, ErrorWrongCommentFormat, []string{}}, "Warn:  path5:5:type /*\\n * \\n */ is suggested"},
        { Record{"path6", 6, ErrorCopyRightDate, []string{"2020-2021", "2020-1-1"}}, "Warn:  path6:6:Please check copy right date 2020-2021 create time 2020-1-1"},

        { Record{"path7", 7, ErrorUsingTab, []string{}}, "Internal Error: param num incorrect"},
        { Record{"path7", 7, ErrorUsingTab, []string{"12", "13"}}, "Internal Error: param num incorrect"},
        { Record{"path7", 7, ErrorInvalid, []string{}}, "Internal Error: type mismatch"},
    }

    for _, test := range(tests) {
        ret := test.r.GetStr()
        if ret != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.r, test.out, ret)
        }
    }
}

func TestLess(t *testing.T) {
    tests := [] struct {
        r1 Record
        r2 Record
        out bool
    } {
        { Record{"path1", 1, 1, []string{}}, Record{string("path2"), 1, 1, []string{}}, true},
        { Record{"path2", 1, 1, []string{}}, Record{string("path1"), 1, 1, []string{}}, false},
        { Record{"path1", 1, 1, []string{}}, Record{string("path1"), 2, 1, []string{}}, true},
        { Record{"path1", 2, 1, []string{}}, Record{string("path1"), 1, 1, []string{}}, false},
        { Record{"path1", 1, 1, []string{}}, Record{string("path1"), 1, 2, []string{}}, true},
        { Record{"path1", 1, 2, []string{}}, Record{string("path1"), 1, 1, []string{}}, false},
        { Record{"path1", 1, 1, []string{}}, Record{string("path1"), 1, 1, []string{"para1"}}, true},
        { Record{"path1", 1, 1, []string{"para1"}}, Record{string("path1"), 1, 1, []string{}}, false},
    }

    for _, test := range(tests) {
        ret := test.r1.Less(test.r2)
        if ret != test.out {
            t.Errorf("test fail, input:%#v %#v expected:%#v output:%#v",
                test.r1, test.r2, test.out, ret)
        }
    }
}

func TestSwap(t *testing.T) {
    r1 := Record{"path1", 1, 1, []string{"para1"}}
    r1_copy := Record{"path1", 1, 1, []string{"para1"}}
    r2 := Record{"path2", 2, 2, []string{"para2"}}
    r2_copy := Record{"path2", 2, 2, []string{"para2"}}

    r1.Swap(&r2)
    if !reflect.DeepEqual(r2, r1_copy) || !reflect.DeepEqual(r1, r2_copy) {
        t.Errorf("test fail")
    }
}