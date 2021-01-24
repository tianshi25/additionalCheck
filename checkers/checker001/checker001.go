package checker001

import (
    . "tianshi25.github.com/additionalCheck/tool"
    . "tianshi25.github.com/additionalCheck/db"
    "strings"
)

const RULE_ID int = 1

var rule = Rule {
    Id:RULE_ID,
    Name:"TabDetected",
    Level:ERROR,
    Fmt:"tab is detected at position %v",
    ParaNum:1,
    Exec:check,
    Info:"tab is detected",
}

func init() {
    RegisterRule(rule)
}

func check(filePath, s string) ([]Record) {
    ret := []Record{}
    s = RemoveWindowsLineEnd(s)
    for i, line := range(strings.Split(s, "\n")) {
        tabPos := strings.Index(line, "\t")
        if (tabPos != -1) {
            r := NewRecord(filePath, i + 1, RULE_ID,
                []interface{}{tabPos})
            ret = append(ret, r)
        }
    }
    return ret
}