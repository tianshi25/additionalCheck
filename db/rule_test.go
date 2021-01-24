package db

import (
    "testing"
    "errors"
    "reflect"
)

func TestRegisterRule(t *testing.T) {
    tests := []struct {
        rule Rule
        len int
    } {
        {
            Rule { Id:1, Name:"1", Level:ERROR,
                Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
            1,
        },
    }

    for _, test := range(tests) {
        RegisterRule(test.rule)
        ret := len(rules)
        if ret != test.len {
            t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
        }
        ClearRules()
    }
}

func TestGetRule(t *testing.T) {
    tests := []struct {
        rules []Rule
        id int
        retId int
        retErr error
    } {
        {
            [] Rule {
                Rule { Id:1, Name:"1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            1,
            1,
            nil,
        },
        {
            [] Rule {
                Rule { Id:1, Name:"1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            3,
            0,
            errors.New("rule id not registered"),
        },
    }

    for _, test := range(tests) {
        for _, rule := range(test.rules) {
            RegisterRule(rule)
        }
        ret, err := GetRule(test.id)
        if ret.Id != test.retId || !reflect.DeepEqual(err, test.retErr) {
            t.Errorf("test fail, test:%#v | ret:%#v err:%#v",
                test , ret, err)
        }
        ClearRules()
    }
}

func TestGetResultStr(t *testing.T) {
    tests := []struct {
        rules []Rule
        id int
        filePath string
        lineNum int
        para []interface{}
        ret string
    } {
        {
            [] Rule {
                Rule { Id:1, Name:"1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            1,
            "path1",
            100,
            []interface{}{1},
            "ERROR   path1:100:str 1",
        },
        {
            [] Rule {
                Rule { Id:1, Name:"1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            1,
            "path1",
            100,
            []interface{}{1, 2},
            INVALID_STR,
        },
        {
            [] Rule {},
            3,
            "path1",
            100,
            []interface{}{1},
            INVALID_STR,
        },
    }

    for _, test := range(tests) {
        for _, rule := range(test.rules) {
            RegisterRule(rule)
        }
        ret := GetResultStr(test.id, test.filePath, test.lineNum, test.para)
        if ret != test.ret {
            t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
        }
        ClearRules()
    }
}

func TestGetInfo(t *testing.T) {
    tests := []struct {
        rules []Rule
        id int
        ret string
    } {
        {
            [] Rule {
                Rule { Id:1, Name:"1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            1,
            "Info 1",
        },
        {
            [] Rule {},
            3,
            INVALID_STR,
        },
    }

    for _, test := range(tests) {
        for _, rule := range(test.rules) {
            RegisterRule(rule)
        }
        ret := GetInfo(test.id)
        if ret != test.ret {
            t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
        }
        ClearRules()
    }
}

func TestGetBrief(t *testing.T) {
    tests := []struct {
        rules []Rule
        ret string
    } {
        {
            [] Rule {
                Rule { Id:1, Name:"Name1", Level:ERROR,
                    Fmt:"str %v", ParaNum:1, Exec:nil, Info:"Info 1"},
                Rule { Id:2, Name:"Name2", Level:ERROR,
                    Fmt:"str %v", ParaNum:2, Exec:nil, Info:"Info 2"},
            },
            "  1 Name1\n  2 Name2",
        },
        {
            [] Rule {},
            "",
        },
    }

    for _, test := range(tests) {
        for _, rule := range(test.rules) {
            RegisterRule(rule)
        }
        ret := GetBriefs()
        if ret != test.ret {
            t.Errorf("test fail, test:%#v | ret:%#v", test, ret)
        }
        ClearRules()
    }
}