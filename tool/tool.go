package tool

import (
    "strings"
    "regexp"
)

func ContainsInt(l []int, t int) bool {
    for _, a := range l {
        if a == t {
            return true
        }
    }
    return false
}

func ContainsString(l []string, t string) bool {
    for _, a := range l {
        if a == t {
            return true
        }
    }
    return false
}

func RemoveWindowsLineEnd(s string) string {
    return strings.Replace(s, "\r\n", "\n", -1)
}

func GetAllComments(s string) []string {
    r := regexp.MustCompile(`(?sm)//.*?$|/\*.*?\*/`)
    ret := r. FindAllString(s, -1)
    if ret != nil {
        return ret;
    }
    return []string{}
}

func getReplacementForComments(comment []byte) []byte {
    lineCount := strings.Count(string(comment), "\n")
    return []byte("" + strings.Repeat("\n", lineCount))
}

func ReplaceComments(s string) string {
    r := regexp.MustCompile(` *((?sm)//.*?$|/\*.*?\*/)`)
    return string(r.ReplaceAllFunc([]byte(s), getReplacementForComments))
}