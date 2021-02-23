package main

import (
	"fmt"
	_ "github.com/tianshi25/additionalCheck/checkers"
	. "github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/filter"
	"github.com/tianshi25/additionalCheck/logs"
	// . "github.com/tianshi25/additionalCheck/filter"
	. "github.com/tianshi25/additionalCheck/para"
	"io/ioutil"
	"strings"
)

// var files, ignoreCheckerIds = GetPar

func main() {
	// files, ignoreCheckerIds, getInfoCheckerId := GetPara()
	files, ignoreCheckerIds, getInfoCheckerId := GetPara()
	if getInfoCheckerId != 0 {
		fmt.Printf("Check %v info: "+GetInfo(getInfoCheckerId)+"\n", getInfoCheckerId)
		return
	}
	records := doJob(files, ignoreCheckerIds)
	fmt.Println(strings.Join(GetRecordsStr(records), "\n"))
}

func doJobInParallel(filePath string, fileContent string, checkId int, ch chan []Record) {
	checkResult := make([]Record, 0)
	rule, err := GetRule(checkId)
	if err != nil {
		logs.E("checker id %v not found", checkId)
		ch <- checkResult
		return
	}
	checkResult = rule.Exec(filePath, fileContent)
	checkResult = filter.FilterRecords(checkResult)
	ch <- checkResult
}

func doJob(fileList []string, checkerIds []int) (ret []Record) {
	count := 0
	ch := make(chan []Record)
	for _, filePath := range fileList {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			logs.E("fail to open file %v", filePath)
		}
		for _, id := range checkerIds {
			count += 1
			go doJobInParallel(filePath, string(content), id, ch)
		}
	}
	for i := 0; i < count; i++ {
		logs.E("tianshi count %d", i)
		ret = append(ret, <-ch...)
	}
	return
}
