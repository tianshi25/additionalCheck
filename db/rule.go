package db

import (
    "tianshi25.github.com/additionalCheck/logs"
    "fmt"
    "errors"
    "strings"
)

const (
    INVALID_STR string = "Internal Error: fail to form string"
)

type CheckFunc func (filePath string, content string) []Record

type Rule struct {
    Id int
    Name string
    Level RuleLevel
    Fmt string
    ParaNum int
    Exec CheckFunc
    Info string
}

func (rule *Rule) getResultStr(filePath string, lineNum int, para []interface{}) string {
    if (len(para) != rule.ParaNum) {
        logs.E("record para num %v is different from expected %v", len(para), rule.ParaNum)
        return INVALID_STR
    }

    localPara := []interface{}{filePath, lineNum}
    localPara = append(localPara, para...)

    return rule.Level.getStr() + fmt.Sprintf(" %v:%v:" + rule.Fmt, localPara...)
}

func (rule *Rule) getBrief() string {
    return fmt.Sprintf("%3v %v", rule.Id, rule.Name)
}

var rules []Rule

func GetRule(id int) (Rule, error) {
    for _, rule := range(rules) {
        if rule.Id == id {
            return rule, nil
        }
    }
    logs.V("rule id %v not registered", id)
    return Rule{}, errors.New("rule id not registered")
}

func GetRules() []Rule {
    return rules
}

func ClearRules() {
    rules = nil
}

func RegisterRule(rule Rule) {
    _, err := GetRule(rule.Id)
    if err == nil {
        logs.E("rule %v already registerd", rule.Id)
        return;
    }
    rules = append(rules, rule)
}

func GetResultStr(id int, filePath string, lineNum int, para []interface{}) string {
    rule, err := GetRule(id)
    if err != nil {
        logs.E("rule %v not registerd", id)
        return INVALID_STR
    }
    return rule.getResultStr(filePath, lineNum, para)
}

func GetInfo(id int) string {
    rule, err := GetRule(id)
    if err != nil {
        logs.E("rule %v not registerd", id)
        return INVALID_STR
    }
    return rule.Info
}

func GetBriefs() string {
    res := []string{}
    for _, rule := range(rules) {
        res = append(res, rule.getBrief())
    }
    return strings.Join(res, "\n")
}