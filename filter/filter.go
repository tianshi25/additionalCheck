package filter

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	. "github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/logs"
	. "github.com/tianshi25/additionalCheck/tool"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var needCareKeys map[string]int
var needCareFiles map[string]int

func init() {
	needCareKeys = make(map[string]int)
	needCareFiles = make(map[string]int)
}

func checkError(err error) {
	if err == nil {
		return
	}

	logs.E("%v occurs during git operation", err)
	os.Exit(1)
}

func SetFilter(path string) {
	repo, err := git.PlainOpen(path)
	checkError(err)

	head, err := repo.Head()
	checkError(err)

	lastCommit, err := repo.CommitObject(head.Hash())
	checkError(err)

	beforeLastCommit, err := lastCommit.Parent(0)
	checkError(err)

	patch, err := lastCommit.Patch(beforeLastCommit)
	checkError(err)

	patchStr := patch.String()
	patchStr = RemoveWindowsLineEnd(patchStr)
	reg := regexp.MustCompile("\n(\\+\\+\\+ b\\/(.*)|@@ \\-\\d+,\\d+ \\+(\\d+),(\\d+) @@)")
	currFilePath := ""
	for _, matched := range reg.FindAllStringSubmatch(patchStr, -1) {
		if len(matched) != 5 {
			logs.E("Invalid match result %v", strings.Join(matched, "|"))
			continue
		}
		matched[0] = strings.Trim(matched[0], "\n")
		if strings.HasPrefix(matched[0], "+++") {
			currFilePath = matched[2]
			needCareFiles[currFilePath] = 1
			addKeyToFilter(currFilePath, 0, 1)
		} else if strings.HasPrefix(matched[0], "@@") {
			start, err := strconv.Atoi(matched[3])
			checkError(err)
			num, err := strconv.Atoi(matched[4])
			checkError(err)
			addKeyToFilter(currFilePath, start, num)
		} else {
			logs.E("processed matched regex error")
		}
	}

}

func GetFileListFromFilter(filePaths []string) (ret []string) {
	if len(needCareFiles) == 0 {
		return filePaths
	}
	ret = make([]string, 0, len(needCareFiles))
	for k := range needCareFiles {
		ret = append(ret, k)
	}
	return
}

func FilterRecords(records []Record) (ret []Record) {
	if len(needCareKeys) == 0 {
		return records
	}
	for _, record := range records {
		if recordPassFilter(record) {
			ret = append(ret, record)
		}
	}
	return
}

func getKeyStr(record Record) string {
	return record.GetPathLineStr()
}

func addKeyToFilter(filePath string, start, num int) {
	// extend check to one line above and below
	context := 1
	for i := -context; i < num + context; i++ {
		lineNum := start + i
		key := fmt.Sprintf("%v:%v", filePath, lineNum)
		needCareKeys[key] = 1
	}
}

func recordPassFilter(record Record) bool {
	// key not in no pass list
	if needCareKeys[getKeyStr(record)] == 0 {
		return false
	}
	return true
}
