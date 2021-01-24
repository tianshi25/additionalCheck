package checker201

import (
	"strings"
	. "tianshi25.github.com/additionalCheck/db"
	. "tianshi25.github.com/additionalCheck/tool"
)

const RULE_ID int = 201

var rule = Rule{
	Id:      RULE_ID,
	Name:    "OperatorMisPos",
	Level:   ERROR,
	Fmt:     "Operator should be located at line %v",
	ParaNum: 1,
	Exec:    check,
	Info:    "For java operator operator shoud be locate at line start. For C++, line end",
}

func init() {
	RegisterRule(rule)
}

func endsWithAny(s string, target []string) bool {
	s = strings.TrimSpace(s)
	for _, t := range target {
		if strings.HasSuffix(s, t) {
			return true
		}
	}
	return false
}

func startsWithAny(s string, target []string) bool {
	s = strings.TrimSpace(s)
	for _, t := range target {
		if strings.HasPrefix(s, t) {
			return true
		}
	}
	return false
}

func check(filePath, s string) []Record {
	ret := []Record{}
	s = RemoveWindowsLineEnd(s)
	s = ReplaceComments(s)

	opEndOfLine := true
	if strings.HasSuffix(filePath, ".c") || strings.HasSuffix(filePath, ".cpp") {
		opEndOfLine = true
	} else if strings.HasSuffix(filePath, ".java") {
		opEndOfLine = false
	} else {
		return ret
	}

	ops := []string{"|", "&", "+", "-", "*"}
	for i, line := range strings.Split(s, "\n") {
		if opEndOfLine && startsWithAny(line, ops) {
			r := NewRecord(filePath, i+1, RULE_ID, []interface{}{"end"})
			ret = append(ret, r)
		}
		if !opEndOfLine && endsWithAny(line, ops) {
			r := NewRecord(filePath, i+1, RULE_ID, []interface{}{"start"})
			ret = append(ret, r)
		}
	}
	return ret
}
