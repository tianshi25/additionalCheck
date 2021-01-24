package main

import (
	"fmt"
	_ "github.com/tianshi25/additionalCheck/checkers"
	. "github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/logs"
	. "github.com/tianshi25/additionalCheck/para"
	. "github.com/tianshi25/additionalCheck/tool"
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

func doJob(fileList []string, ignoreCheckerId []int) (ret []Record) {
	for _, file := range fileList {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			logs.E("fail to open file %v", file)
		}
		s := string(content)
		for _, rule := range GetRules() {
			if ContainsInt(ignoreCheckerId, rule.Id) {
				continue
			}
			checkResult := rule.Exec(file, s)
			ret = append(ret, checkResult...)
		}
	}
	return
}
