package checker102

import (
	. "github.com/tianshi25/additionalCheck/db"
	. "github.com/tianshi25/additionalCheck/tool"
	"regexp"
	"strings"
)

const RULE_ID int = 102

var rule = Rule{
	Id:      RULE_ID,
	Name:    "MultilineCommitFormatWrong",
	Level:   SUGGEST,
	Fmt:     "type /*\\n * \\n */ multiline commit is suggested",
	ParaNum: 0,
	Exec:    check,
	Info:    "type /*\\n * \\n */ multiline commit is suggested",
}

func init() {
	RegisterRule(rule)
}

func check(filePath, s string) []Record {
	var ret []Record
	s = RemoveWindowsLineEnd(s)
	comments, lineNums := GetAllCommentsWithLineNum(s)
	for i, comment := range comments {
		if !strings.HasPrefix(comment, "/*") {
			continue
		}
		if strings.Count(comment, "\n") == 0 {
			continue
		}
		r := regexp.MustCompile(`/\*\n( +\* .*\n)+ *\*/`)
		match := r.FindAllString(s, -1)
		if match == nil {
			r := NewRecord(filePath, lineNums[i], RULE_ID, []interface{}{})
			ret = append(ret, r)
		}
	}
	return ret
}
