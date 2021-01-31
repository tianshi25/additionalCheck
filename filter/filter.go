package filter

import (
	"fmt"
	. "github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/logs"
	"github.com/go-git/go-git/v5"
)

var noPassKeys map[string]int

func SetFilter(path string) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		logs.W("%v is not git repo", path)
		return
	}

	head, err := repo.Head()
	if err != nil {
		logs.E("fail to get git HEAD")
		return
	}

	fmt.Printf("%#v", head)
}

func filterRecords(records []Record) (ret []Record){
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

func recordPassFilter(record Record) bool {
	// key not in no pass list
	if noPassKeys[getKeyStr(record)] == 0 {
		return false
	}
	return true
}