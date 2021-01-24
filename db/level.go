package db

type RuleLevel int

const (
	ERROR   RuleLevel = 0
	SUGGEST RuleLevel = 1
	NOTE    RuleLevel = 2
)

func (level RuleLevel) getStr() string {
	if level == ERROR {
		return "ERROR  "
	}
	if level == SUGGEST {
		return "SUGGEST"
	}
	return "NOTE   "
}
