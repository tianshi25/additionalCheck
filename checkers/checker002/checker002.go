package checker002

import (
	. "github.com/tianshi25/additionalCheck/db"
	. "github.com/tianshi25/additionalCheck/tool"
	"strings"
)

const RULE_ID int = 2

var rule = Rule{
	Id:      RULE_ID,
	Name:    "IndentMultiplesOfFour",
	Level:   SUGGEST,
	Fmt:     "indent should be multiples of 4 and increase by 4",
	ParaNum: 0,
	Exec:    check,
	Info:    "indent should be multiples of 4 and increase by 4",
}

func init() {
	RegisterRule(rule)
}

func check(filePath, s string) []Record {
	ret := []Record{}
	s = RemoveWindowsLineEnd(s)
	s = ReplaceComments(s)
	lastIndent := 0
	for i, line := range strings.Split(s, "\n") {
		currIndent := GetIndentSpaceNum(line)
		if currIndent%4 != 0 ||
			(currIndent-lastIndent > 4) {
			r := NewRecord(filePath, i+1, RULE_ID, []interface{}{})
			ret = append(ret, r)
		}
		lastIndent = (currIndent / 4) * 4
	}
	return ret
}
