package checker

import (
    "testing"
    "reflect"
    "strings"
)

func TestGetIndentSpaceNum(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "    abc d", 4 },
        { "    \n", 4 },
        { "abc    de", 0 },
        { "", 0 },
    }

    for _, test := range(tests) {
        num := getIndentSpaceNum(test.in)
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorWrongIndent(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "     abc\n        efg\n", 1 },
        { "    abc\n}efg\n", 0 },
        { "    abc\n        efg\n", 0 },
        { "    abc\n            efg\n", 1 },
        { "    abc\n         efg\n", 1 },
        { "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorWrongIndent("file1", test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorExtraSpace(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "   ", 1 },
        { "   \n", 1 },
        { "    abc", 0 },
        { "    abc ", 1 },
        { "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorExtraSpace("file1", test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorMixComment(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "//line1\n//line2", 0 },
        { "/*comm\nen\nt1*///line2", 1 },
        { "//line2\n/*comm\nen\nt1*/", 1 },
        { "/*comm\nen\nt1*//*comm\nen\nt2*/", 0 },
        { "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorMixComment("file1", test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorWrongCommentFormat(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "/*\n * comment1\n * comment2\n */", 0 },
        { "    /*\n     * comment1\n     * comment2\n */", 0 },
        { "    /* comment2 */", 1 },
        { "    /* \ncomment2 */", 1 },
        { "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorWrongCommentFormat("file1", test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorCopyRightDate(t *testing.T) {
    tests := [] struct {
        in string
        out int
    } {
        { "-2021", 0 },
        { "/* 1993-2021 */", 0 },
        { "/* 1993-2020 */", 1 },
        { "-2020", 0 },
        { "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorCopyRightDate("file1", test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}

func TestCheckErrorLogicOpMisPosition(t *testing.T) {
    tests := [] struct {
        filePath string
        in string
        out int
    } {
        { "1.c", "line", 0 },
        { "1.c", "line ||", 0 },
        { "1.c", "line &&", 0 },
        { "1.c", "  || line", 1 },
        { "1.c", "  && line", 1 },
        { "1.java", "line", 0 },
        { "1.java", "line ||", 1 },
        { "1.java", "line &&", 1 },
        { "1.java", "  || line", 0 },
        { "1.java", "  && line", 0 },
        { "1.c", "", 0 },
    }

    for _, test := range(tests) {
        db := CheckErrorLogicOpMisPosition(test.filePath, test.in)
        num := strings.Count(db.GetStr(), "\n")
        if num != test.out {
            t.Errorf("test fail, input:%#v expected:%#v output:%#v",
                test.in, test.out, num)
        }
    }
}