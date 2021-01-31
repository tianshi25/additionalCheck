package filter

import (
	"os"
	. "github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/logs"
	"github.com/go-git/go-git/v5"
)

var noPassKeys map[string]int

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

	logs.E("%v %v", lastCommit, lastCommit.NumParents())
    beforeLastCommit, err := lastCommit.Parent(0)
    checkError(err)

    patch, err := lastCommit.Patch(beforeLastCommit)
    checkError(err)

    logs.E("%v", patch)
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