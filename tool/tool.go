package tool

import (
	"github.com/tianshi25/additionalCheck/logs"
	"path/filepath"
	"regexp"
	"strings"
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

func GetAllCommentsWithLineNum(s string) (sRet []string, iRet []int) {
	r := regexp.MustCompile(`(?sm)//.*?$|/\*.*?\*/`)
	indexRanges := r.FindAllStringIndex(s, -1)
	if indexRanges == nil {
		return []string{}, []int{}
	}

	for _, indexRange := range indexRanges {
		comment := s[indexRange[0]:indexRange[1]]
		lineNum := strings.Count(s[:indexRange[0]], "\n") + 1
		sRet = append(sRet, comment)
		iRet = append(iRet, lineNum)
	}
	return sRet, iRet
}

func GetAllComments(s string) []string {
	r := regexp.MustCompile(`(?sm)//.*?$|/\*.*?\*/`)
	ret := r.FindAllString(s, -1)
	if ret != nil {
		return ret
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

func FileExtensionIsC(path string) bool {
	extensions := []string{"c", "cpp", "h", "hpp"}
	return fileExtensionMatchReg(path, extensions)
}

func FileExtensionIsJava(path string) bool {
	extensions := []string{"java"}
	return fileExtensionMatchReg(path, extensions)
}

func fileExtensionMatchReg(currPath string, extensions []string) bool {
	regStr := `.+\.(` + strings.Join(extensions, "|") + `)$`
	if matched, err := regexp.MatchString(regStr, filepath.Base(currPath)); err != nil {
		logs.E("error occur when match string")
		return false
	} else if matched {
		return true
	}
	return false
}

func ConvertWinPath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}