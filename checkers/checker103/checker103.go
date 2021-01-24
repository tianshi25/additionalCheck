package checker103

import (
    . "tianshi25.github.com/additionalCheck/tool"
    . "tianshi25.github.com/additionalCheck/db"
    "regexp"
    "strconv"
    "time"
)

const RULE_ID int = 103

var rule = Rule {
    Id:RULE_ID,
    Name:"CopyRightDateNotUpdated",
    Level:ERROR,
    Fmt:"Please check copy right last modify date",
    ParaNum:0,
    Exec:check,
    Info:"Please check copy right last modify date",
}

func init() {
    RegisterRule(rule)
}

func check(filePath, s string) ([]Record) {
    ret := []Record{}
    s = RemoveWindowsLineEnd(s)
    comments := GetAllComments(s)
    if (len(comments) == 0) {
        return ret
    }
    r := regexp.MustCompile(`-\d\d\d\d`)
    endDate := r.FindAllString(comments[0], -1)
    if endDate == nil {
        return ret
    }

    currDate := "-" + strconv.Itoa(time.Now().Year())
    if endDate[0] != currDate {
        r := NewRecord(filePath, 0, RULE_ID, []interface{}{})
        ret = append(ret, r)
    }
    return ret
}