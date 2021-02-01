package db

import (
	"fmt"
	. "github.com/tianshi25/additionalCheck/tool"
)

type Record struct {
	filePath  string
	lineNum   int
	errorId   int
	errorPara []interface{}
}

func NewRecord(filePath string, lineNum int, errorId int, errorPara []interface{}) Record {
	return Record{filePath, lineNum, errorId, errorPara}
}

func (r1 Record) GetStr() string {
	return GetResultStr(r1.errorId, r1.filePath, r1.lineNum, r1.errorPara)
}

func GetRecordsStr(l []Record) (ret []string) {
	for _, r := range l {
		ret = append(ret, r.GetStr())
	}
	return
}

func (r1 Record) Less(r2 Record) bool {
	if r1.filePath < r2.filePath {
		return true
	}
	if r1.filePath > r2.filePath {
		return false
	}

	if r1.lineNum < r2.lineNum {
		return true
	}
	if r1.lineNum > r2.lineNum {
		return false
	}

	if r1.errorId < r2.errorId {
		return true
	}
	if r1.errorId > r2.errorId {
		return false
	}

	return len(r1.errorPara) < len(r2.errorPara)
}

func (r1 *Record) Swap(r2 *Record) {
	r1.filePath, r2.filePath = r2.filePath, r1.filePath
	r1.lineNum, r2.lineNum = r2.lineNum, r1.lineNum
	r1.errorId, r2.errorId = r2.errorId, r1.errorId
	r1.errorPara, r2.errorPara = r2.errorPara, r1.errorPara
}

func (r1 *Record) GetPathLineStr() string {
	return ConvertWinPath(fmt.Sprintf("%v:%v", r1.filePath, r1.lineNum))
}
