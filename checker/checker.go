package checker

import (
    "regexp"
    "strings"
    "tianshi25.github.com/additionalCheck/recordDb"
    "strconv"
    "time"
)

type ErrorType int

const (
    // white space
    ErrorWrongIndent ErrorType = 2
    ErrorExtraSpace ErrorType = 3

    // comment
    ErrorMixComment ErrorType = 11
    ErrorWrongCommentFormat ErrorType = 12
    ErrorCopyRightDate ErrorType = 13

    // other
    ErrorLogicOpMisPosition ErrorType = 21

    ErrorInvalid ErrorType = 100
)

var ErrorTemplate  = [] struct {
    id ErrorType
    level int
    template string
    paraNum int
} {
    {ErrorWrongIndent, logs.Warn, "indent should increase 0 or 4", 0},
    {ErrorExtraSpace, logs.Warn, "extra space is detected at position %v", 1},

    {ErrorMixComment, logs.Warn, "mixing two comment type", 0},
    {ErrorWrongCommentFormat, logs.Warn, "type /*\\n * \\n */ is suggested", 0},
    {ErrorCopyRightDate, logs.Warn, "Please check copy right last modify date", 0},

    {ErrorLogicOpMisPosition, logs.Error, "logic Operater: java line shart; C++ line end", 0},
}



// ErrorUsingTab ErrorType = 1
func CheckErrorUsingTab(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    db := recordDb.RecordDb{}
    for i, line := range(strings.Split(s, "\n")) {
        tabPos := strings.Index(line, "\t")
        if (tabPos != -1) {
            r := record.NewRecord(filePath, i + 1, record.ErrorUsingTab,
                []string{strconv.Itoa(tabPos)})
            db.Append(r)
        }
    }
    return db
}

func getIndentSpaceNum(s string) int {
    count := 0
    for _, c := range(s) {
        if c != ' ' {
            break;
        }
        count += 1
    }
    return count
}

// ErrorWrongIndent ErrorType = 2
func CheckErrorWrongIndent(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    s = ReplaceComments(s)
    db := recordDb.RecordDb{}
    lastIndent := 0
    for i, line := range(strings.Split(s, "\n")) {
        currIndent := getIndentSpaceNum(line)
        if currIndent % 4 != 0 ||
            (currIndent - lastIndent > 4) {
            r := record.NewRecord(filePath, i + 1, record.ErrorWrongIndent,
                []string{})
            db.Append(r)
        }
        lastIndent = currIndent;
    }
    return db
}
// ErrorExtraSpace ErrorType = 3
func CheckErrorExtraSpace(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    s = ReplaceComments(s)
    db := recordDb.RecordDb{}
    for i, line := range(strings.Split(s, "\n")) {
        indent := getIndentSpaceNum(line)
        if len(line) > 0 && line[len(line) - 1] == ' ' {
            r := record.NewRecord(filePath, i + 1, record.ErrorExtraSpace,
                []string{"0"})
            db.Append(r)
            continue
        }

        pos := strings.Index(line[indent:], "  ")
        if pos != -1 {
            r := record.NewRecord(filePath, i + 1, record.ErrorExtraSpace,
                []string{strconv.Itoa(pos)})
            db.Append(r)
            continue
        }

    }
    return db
}

// ErrorMixComment ErrorType = 11
func CheckErrorMixComment(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    comments := getAllComments(s)
    db := recordDb.RecordDb{}
    singleLine := false
    multiLine := false
    for _, comment := range(comments) {
        if strings.HasPrefix(comment, "//") {
            singleLine = true
        }
        if strings.HasPrefix(comment, "/*") {
            multiLine = true
        }
    }
    if singleLine && multiLine {
        r := record.NewRecord(filePath, 0, record.ErrorMixComment,
            []string{})
        db.Append(r)
    }

    return db
}

// ErrorWrongCommentFormat ErrorType = 12
func CheckErrorWrongCommentFormat(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    comments := getAllComments(s)
    db := recordDb.RecordDb{}
    for i, comment := range(comments) {
        if !strings.HasPrefix(comment, "/*") {
            continue
        }
        r := regexp.MustCompile(`/\*\n( +\* .*\n)+ *\*/`)
        ret := r. FindAllString(s, -1)
        if ret == nil {
            r := record.NewRecord(filePath, i + 1, record.ErrorMixComment, []string{})
            db.Append(r)
        }
    }

    return db
}

// ErrorCopyRightDate ErrorType = 13
func CheckErrorCopyRightDate(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    comments := getAllComments(s)

    db := recordDb.RecordDb{}
    if (len(comments) == 0) {
        return db
    }

    r := regexp.MustCompile(`-\d\d\d\d`)
    endDate := r. FindAllString(comments[0], -1)
    if endDate == nil {
        return db
    }

    currDate := "-" + strconv.Itoa(time.Now().Year())
    if endDate[0] != currDate {
        r := record.NewRecord(filePath, 0, record.ErrorCopyRightDate, []string{})
        db.Append(r)
    }

    return db
}

// ErrorLogicOpMisPosition ErrorType = 21
func CheckErrorLogicOpMisPosition(filePath, s string) recordDb.RecordDb {
    s = removeWindowsLineEnd(s)
    s = ReplaceComments(s)

    db := recordDb.RecordDb{}

    opEndOfLine := true
    if (strings.HasSuffix(filePath, ".c") || strings.HasSuffix(filePath, ".cpp")) {
        opEndOfLine = true
    } else if (strings.HasSuffix(filePath, ".java")) {
        opEndOfLine = false
    } else {
        return db
    }

    for i, line := range(strings.Split(s, "\n")) {
        if (opEndOfLine) {
            indent := getIndentSpaceNum(line)
            if (strings.HasPrefix(line[indent:], "||") || strings.HasPrefix(line[indent:], "&&")) {
                r := record.NewRecord(filePath, i+1, record.ErrorLogicOpMisPosition, []string{})
                db.Append(r)
            }
        } else {
            if (strings.HasSuffix(line, "||") || strings.HasSuffix(line, "&&")) {
                r := record.NewRecord(filePath, i+1, record.ErrorLogicOpMisPosition, []string{})
                db.Append(r)
            }
        }
    }

    return db
}

var CheckList = [] struct {
    ErrorId record.ErrorType
    F func(filePath, s string) recordDb.RecordDb
} {
    {record.ErrorUsingTab, CheckErrorUsingTab},
    {record.ErrorWrongIndent, CheckErrorWrongIndent},
    {record.ErrorExtraSpace, CheckErrorExtraSpace},

    {record.ErrorMixComment, CheckErrorMixComment},
    {record.ErrorWrongCommentFormat, CheckErrorWrongCommentFormat},
    {record.ErrorCopyRightDate, CheckErrorCopyRightDate},

    {record.ErrorLogicOpMisPosition, CheckErrorLogicOpMisPosition},
}