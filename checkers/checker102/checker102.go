package checker102

import (
    . "tianshi25.github.com/additionalCheck/tool"
    . "tianshi25.github.com/additionalCheck/db"
    "strings"
    "regexp"
)

const RULE_ID int = 102

var rule = Rule {
    Id:RULE_ID,
    Name:"MultilineCommitFormatWrong",
    Level:SUGGEST,
    Fmt:"type /*\\n * \\n */ multiline commit is suggested",
    ParaNum:0,
    Exec:check,
    Info:"type /*\\n * \\n */ multiline commit is suggested",
}

func init() {
    RegisterRule(rule)
}

func check(filePath, s string) ([]Record) {
    ret := []Record{}
    s = RemoveWindowsLineEnd(s)
    comments := GetAllComments(s)
    for i, comment := range(comments) {
        if !strings.HasPrefix(comment, "/*") {
            continue
        }
        r := regexp.MustCompile(`/\*\n( +\* .*\n)+ *\*/`)
        match := r. FindAllString(s, -1)
        if match == nil {
            r := NewRecord(filePath, i + 1, RULE_ID, []interface{}{})
            ret = append(ret, r)
        }
    }
    return ret
}