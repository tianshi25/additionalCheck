package tool

import (
    "strings"
    "regexp"
)

func removeWindowsLineEnd(s string) string {
    return strings.Replace(s, "\r\n", "\n", -1)
}

func getAllComments(s string) []string {
    r := regexp.MustCompile(`(?sm)//.*?$|/\*.*?\*/`)
    ret := r. FindAllString(s, -1)
    if ret != nil {
        return ret;
    }
    return []string{}
}

func getReplacementForComments(comment string) string {
    lineCount := strings.Count(comment, "\n")
    return "" + strings.Repeat("\n", lineCount)
}

func ReplaceComments(s string) string {
    allComments := getAllComments(s)
    for _, comment := range(allComments) {
        s = strings.ReplaceAll(s, comment, getReplacementForComments(comment));
    }
    return s;
}