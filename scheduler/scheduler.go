package scheduler

import (
    "tianshi25.github.com/checker"
    "tianshi25.github.com/record"
    "tianshi25.github.com/recordDb"
    "strings"
    "ioutil"
    "logs"
)

func contains(l []record.ErrorType, t record.ErrorType) bool {
    for _, a := range l {
        if a == t {
            return true
        }
    }
    return false
}

func DoJob(fileList []string, skipCheckList []record.ErrorType) recordDb.RecordDb {
    db := recordDb.RecordDb{}
    for _, file := range(fileList) {
        if !strings.HasSuffix(file, ".java") &&
            !strings.HasSuffix(file, ".cpp") &&
            !strings.HasSuffix(file, ".c") {
            continue;
        }
        content, err := ioutil.ReadFile("golangcode.txt")
        if err != nil {
            logs.E("fail to open file %v", file)
        }
        s := string(content)
        for _, check := range(checker.CheckList) {
            if contains(skipCheckList, check.ErrorId) {
                continue
            }
            db.Concate(check.F(file, s))
        }
    }
}