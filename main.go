package main

import (
    "flag"
    "tianshi25.github.com/record"
    "tianshi25.github.com/recordDb"
    "tianshi25.github.com/logs"
    "tianshi25.github.com/checker"
    "strconv"
    "os"
    "path/filepath"
    "strings"
    "io/ioutil"
    "fmt"
)

var ignoreHelp=`ignore check
example: 1,11,21
ErrorUsingTab             1
ErrorWrongIndent          2
ErrorExtraSpace           3
ErrorMixComment           11
ErrorWrongCommentFormat   12
ErrorCopyRightDate        13
ErrorLogicOpMisPosition   21`

var path = flag.String("path", ".", "file or folder path")
var ingnoreFlagsStr = flag.String("ignore", "", ignoreHelp)

func main() {
    flag.Parse()

    logs.I("path %#v flags %#v", *path, *ingnoreFlagsStr)
    paths, err := getAllFiles(*path)
    if err != nil {
        return
    }

    flags, err := getAllFlags(*ingnoreFlagsStr)
    if err != nil {
        return
    }

    db := DoJob(paths, flags)
    fmt.Println(db.GetStr())
}

func fileNameIsCOrJava(path string) bool {
    return strings.HasSuffix(path, ".cpp") || strings.HasSuffix(path, ".c") ||
        strings.HasSuffix(path, ".java") || strings.HasSuffix(path, ".go")
}

func getAllFiles(path string) ([]string, error) {
    paths := []string{};
    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            logs.E("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return nil
        }
        if info.IsDir() {
            logs.V("skipping a dir without errors: %+v \n", info.Name())
            return nil
        }
        if !fileNameIsCOrJava(path) {
            logs.V("skipping a file without errors: %+v \n", info.Name())
            return nil
        }
        logs.V("visited file or dir: %q\n", path)
        paths = append(paths, path)
        return nil
    })
    if err != nil {
        logs.E("error walking the path %q: %v\n", path, err)
        return paths, err
    }
    logs.I("files to check:")
    for _, p := range(paths) {
        logs.I(p)
    }
    logs.I("")
    return paths, err
}

func getAllFlags(flagStr string) (flags []record.ErrorType, err error) {
    flags = []record.ErrorType{}
    if (len(flagStr) != 0) {
        for _, c := range(strings.Split(flagStr, ",")) {
            v, err := strconv.Atoi(c)
            if err != nil {
                logs.E("atoi string %v error %v\n", v, err)
                return flags, err
            }
            flags = append(flags, record.ErrorType(v))
        }
    }
    logs.I("flags:")
    for _, f := range(flags) {
        logs.I("%v", f)
    }
    logs.I("")
    return flags, err
}

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
        content, err := ioutil.ReadFile(file)
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
    return db
}