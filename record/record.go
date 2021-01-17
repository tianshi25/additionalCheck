package record

import (
    "tianshi25.github.com/logs"
    "strconv"
)

type ErrorType int

const (
    // white space
    ErrorUsingTab ErrorType = 1
    ErrorWrongIndent ErrorType = 2
    ErrorExtraSpace ErrorType = 3

    // comment
    ErrorMixComment ErrorType = 11
    ErrorWrongCommentFormat ErrorType = 12
    ErrorCopyRightDate ErrorType = 13

    // other
    ErrorLogicOpMisPosition ErrorType = 21

    ErrorInvalid ErrorType = 100
)

var errorTemplate  = [] struct {
    id ErrorType
    level int
    template string
    paraNum int
} {
    {ErrorUsingTab, logs.Error, "tab is detected at position %v", 1},
    {ErrorWrongIndent, logs.Warn, "indent should increase 0 or 4", 0},
    {ErrorExtraSpace, logs.Warn, "extra space is detected at position %v", 1},

    {ErrorMixComment, logs.Warn, "mixing two comment type", 0},
    {ErrorWrongCommentFormat, logs.Warn, "type /*\\n * \\n */ is suggested", 0},
    {ErrorCopyRightDate, logs.Warn, "Please check copy right last modify date", 0},

    {ErrorLogicOpMisPosition, logs.Error, "logic Operater: java line shart; C++ line end", 0},
}

type Record struct {
    filePath string
    lineNum int
    errorType ErrorType
    errorPara []string
}

func NewRecord(filePath string, lineNum int, errorType ErrorType, errorPara []string) Record {
    return Record{filePath, lineNum, errorType, errorPara}
}

func (r Record) GetStr() string {
    for _, t := range(errorTemplate) {
        if (t.id == r.errorType) {
            if (len(r.errorPara) != t.paraNum) {
                logs.E("record %#v para num less then expected", r)
                return "Internal Error: param num incorrect"
            }
            para := []string{r.filePath, strconv.Itoa(r.lineNum)}
            para = append(para, r.errorPara...)
            printPara := make([]interface{}, len(para))
            for i:=0; i<len(para); i++ {
                printPara[i] = para[i]
            }
            return logs.RecordSprintf((int)(t.level), "%v:%v:" + t.template, printPara...)
        }
    }
    logs.E("record %#v type mismatch", r)
    return "Internal Error: type mismatch"
}

func (r1 Record) Less(r2 Record) bool {
    if r1.filePath < r2.filePath {
        return true
    }
    if r1.filePath > r2.filePath {
        return false
    }

    if r1.lineNum < r2.lineNum {
        return true
    }
    if r1.lineNum > r2.lineNum {
        return false
    }

    if r1.errorType < r2.errorType {
        return true
    }
    if r1.errorType > r2.errorType {
        return false
    }

    return len(r1.errorPara) < len(r2.errorPara)
}

func (r1 *Record) Swap(r2 *Record) {
    r1.filePath, r2.filePath = r2.filePath, r1.filePath
    r1.lineNum, r2.lineNum = r2.lineNum, r1.lineNum
    r1.errorType, r2.errorType = r2.errorType, r1.errorType
    r1.errorPara, r2.errorPara = r2.errorPara, r1.errorPara
}