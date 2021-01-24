package tool

import (
    "strings"
    "regexp"
    "path/filepath"
    "tianshi25.github.com/additionalCheck/logs"
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
    r := regexp.MustCompile(`(?sm) *(//.*?$|/\*.*?\*/)`)
    return string(r.ReplaceAllFunc([]byte(s), getReplacementForComments))
}

func GetIndentSpaceNum(s string) int {
    return len(s) - len(strings.TrimLeft(s, " "))
}

func FileExtensionIsC(currPath string) bool {
    extensions := []string{"c", "cpp", "h", "hpp"}
    regStr := `.+\.(` + strings.Join(extensions, "|") + `)$`
    if matched, err := regexp.MatchString(regStr, filepath.Base(currPath)); err != nil {
        logs.E("error occur when match string")
        return false
    } else if matched {
        return true
    }
    return false
}

func FileExtensionIsJava(fileName string) bool {
    return strings.HasSuffix(fileName, ".java") && len(fileName) > len(".java")
}