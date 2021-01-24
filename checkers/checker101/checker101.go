package checker101

import (
	"strings"
	. "tianshi25.github.com/additionalCheck/db"
	. "tianshi25.github.com/additionalCheck/tool"
)

const RULE_ID int = 101

var rule = Rule{
	Id:      RULE_ID,
	Name:    "MixCommentType",
	Level:   NOTE,
	Fmt:     "mixing two comment type",
	ParaNum: 0,
	Exec:    check,
	Info:    "mixing two comment type",
}

func init() {
	RegisterRule(rule)
}

func check(filePath, s string) []Record {
	ret := []Record{}
	s = RemoveWindowsLineEnd(s)
	comments := GetAllComments(s)
	singleLine := false
	multiLine := false
	for _, comment := range comments {
		if strings.HasPrefix(comment, "//") {
			singleLine = true
		}
		if strings.HasPrefix(comment, "/*") {
			multiLine = true
		}
	}
	if singleLine && multiLine {
		record := NewRecord(filePath, 0, RULE_ID, []interface{}{})
		ret = append(ret, record)
	}

	return ret
}
