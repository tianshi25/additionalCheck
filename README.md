# AdditionalCheck

AdditionalCheck is coding style check framework. 

## Usages

```bash
❯ /out/additionalCheck_darwin_amd64 -h
Usage of additionalCheck_darwin_amd64:
Usage of ./out/additionalCheck_darwin_amd64:
  -ext string
    	file extension to check
    	default: c,cpp,h,hpp,java,go (default "c,cpp,h,hpp,java,go")
  -info int
    	get checker id info
    	  1 TabDetected
        ...
  -isGitRepo
    	path is git repo (default true)
  -log string
    	log level
    	value: Error Warn Info Verbose
    	default:Error (default "Error")
  -no string
    	ignore checker id:
    	example: 1,2
    	  1 TabDetected
        ...
  -only string
    	checker id
    	example: 1,2
    	  1 TabDetected
        ...
  -path string
    	paths to check
    	default:. (default ".")
```

# Supported Checkers

```bash
❯ ./out/additionalCheck_darwin_amd64 -info 1
Check 1 info: tab is detected
❯ ./out/additionalCheck_darwin_amd64 -info 2
Check 2 info: indent should be multiples of 4 and increase by 4
❯ ./out/additionalCheck_darwin_amd64 -info 3
Check 3 info: extra space is detected
❯ ./out/additionalCheck_darwin_amd64 -info 101
Check 101 info: mixing two comment type
❯ ./out/additionalCheck_darwin_amd64 -info 102
Check 102 info: type /*\n * \n */ multiline commit is suggested
❯ ./out/additionalCheck_darwin_amd64 -info 103
Check 103 info: Please check copy right last modify date
❯ ./out/additionalCheck_darwin_amd64 -info 201
Check 201 info: For java operator operator shoud be locate at line start. For C++, line end
```

## Project Structure

```bash
.
├── checkers
│   ├── checker001 # checkerxxx contains a checker and its test cases
│   │   ├── checker001.go
│   │   └── checker001_test.go
│   ├── checker...
│   └── checkers.go # enabled checkers
├── db
│   ├── record.go # detected error data structure
│   └── rule.go # checker rule data structure
├── filter
│   └── filter.go # get patch of lastest commit, use this info to filter detected errors
├── logs # colorful log function
├── para
│   └── para.go # command parameters
└── tool
│   └── tool.go # tool functions
├── release.sh # release binary (win mac linux) to ./out
└── main.go # process entry
```

## Bulid Guide

### Prerequisite

go 1.11 and above (need go mod)

### Build script

```bash
./release.sh
```

### Debugging Exceuate Command
```bash
go run .
```



## Checker Example

### Checker001.go

```go
package checker001

import (
	. "github.com/tianshi25/additionalCheck/db"
	. "github.com/tianshi25/additionalCheck/tool"
	"strings"
)

const RULE_ID int = 1 // define rule id

var rule = Rule{ // fill rule info
	Id:      RULE_ID,
	Name:    "TabDetected", // rule name
  Level:   ERROR, // rule level: ERROR, SUGGEST, NOTE
	Fmt:     "tab is detected at position %v", // rule formatting str
	ParaNum: 1, // rule formatting str para num
	Exec:    check, // check function
	Info:    "tab is detected", // -info RULE_ID str
}

func init() {
	RegisterRule(rule) // register rule
}

// filePath: relative path of a file need check expample a/b/c.java
// s: file content
func check(filePath, s string) []Record {
	ret := []Record{}
	s = RemoveWindowsLineEnd(s) // call tool function to remove \r line ending
	for i, line := range strings.Split(s, "\n") { // split content to lines
		tabPos := strings.Index(line, "\t")
		if tabPos != -1 {
			r := NewRecord(filePath, i+1, RULE_ID, // create error record
				[]interface{}{tabPos})
			ret = append(ret, r)
		}
	}
	return ret
}
```

### checkers.go

```go
// add black import of check001. check001 init() will be called in progrom init phase
import _ "github.com/tianshi25/additionalCheck/checkers/checker001"
```

