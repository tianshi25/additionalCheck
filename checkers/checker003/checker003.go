package checker003

import (
	. "github.com/tianshi25/additionalCheck/db"
	. "github.com/tianshi25/additionalCheck/tool"
	"strings"
)

const RULE_ID int = 3

var rule = Rule{
	Id:      RULE_ID,
	Name:    "ExtraSpace",
	Level:   ERROR,
	Fmt:     "extra space is detected at position %v",
	ParaNum: 1,
	Exec:    check,
	Info:    "extra space is detected",
}

func init() {
	RegisterRule(rule)
}

func check(filePath, s string) []Record {
	ret := []Record{}
	s = RemoveWindowsLineEnd(s)

	for i, line := range strings.Split(s, "\n") {
		if strings.HasSuffix(line, " ") {
			record := NewRecord(filePath, i+1, RULE_ID, []interface{}{len(line) - 1})
			ret = append(ret, record)
		}
	}

	s = ReplaceComments(s)
	for i, line := range strings.Split(s, "\n") {

		indent := GetIndentSpaceNum(line)
		pos := strings.Index(line[indent:], "  ")
		if pos != -1 {
			record := NewRecord(filePath, i+1, RULE_ID, []interface{}{indent + pos})
			ret = append(ret, record)
		}
	}
	return ret
}
