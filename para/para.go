package para

import (
	"flag"
	"github.com/tianshi25/additionalCheck/db"
	. "github.com/tianshi25/additionalCheck/filter"
	"github.com/tianshi25/additionalCheck/logs"
	. "github.com/tianshi25/additionalCheck/tool"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetPara() (fileList []string, checkerIds []int, getInfoCheckerId int) {
	searchPath := flag.String("path", ".", "paths to check")
	ignoreStr := flag.String("no", "", "ignore checker id:\nexample: 1,2"+db.GetBriefs())
	onlyStr := flag.String("only", "", "checker id\nexample: 1,2\n"+db.GetBriefs())
	extensionStr := flag.String("ext", "c,cpp,h,hpp,java,go", "file extension to check")
	logLevel := flag.String("log", "Error", "log level\nvalue: Error Warn Info Verbose")
	getInfoCheckerIdAddr := flag.Int("info", 0, "get checker id info\n"+db.GetBriefs())
	isGitRepo := flag.Bool("gitrepo", false,
		`path is git repo
When this para is set, results are filtered by changes in last commit`)

	flag.Parse()
	setLogLevel(*logLevel)
	if len(*ignoreStr) != 0 && len(*onlyStr) != 0 {
		logs.E("ignore and only are conflict para")
		return
	}

	getInfoCheckerId = *getInfoCheckerIdAddr
	extensions := getExtensionList(*extensionStr)

	if len(*onlyStr) != 0 {
		checkerIds = getCheckerId(*onlyStr)
	} else {
		checkerIds = getCheckIdNotIgnore(*ignoreStr)
	}

	err := os.Chdir(*searchPath)
	if err != nil {
		logs.E("fails to cd " + *searchPath)
		return
	}

	if *isGitRepo {
		SetFilter(".")
		fileList = GetFileListFromFilter(fileList)
	} else {
		logs.E("tianshi52")
		fileList = getFileList(".", extensions)
	}
	logs.I("files to check:\n" + strings.Join(fileList, "\n"))
	return
}

func setLogLevel(level string) {
	logLevel := logs.Invalid
	if level == "Error" {
		logLevel = logs.Error
	} else if level == "Warn" {
		logLevel = logs.Warn
	} else if level == "Info" {
		logLevel = logs.Info
	} else if level == "Verbose" {
		logLevel = logs.Verbose
	}

	logs.SetLevel(logLevel)
}

func getCheckerId(flagStr string) []int {
	var ids []int
	for _, s := range strings.Split(flagStr, ",") {
		if len(s) == 0 {
			continue
		}
		id, err := strconv.Atoi(s)
		if err != nil {
			logs.E("atoi string %v error %v\n", id, err)
			continue
		}
		ids = append(ids, id)
	}
	logs.I("ignore ids: %#v", ids)
	return ids
}

var VALID_EXTENSIONS = map[string]int{"c": 1, "h": 1, "cpp": 1, "hpp": 1, "java": 1, "go": 1}

func getExtensionList(extensionStr string) (exts []string) {
	for _, s := range strings.Split(extensionStr, ",") {
		if len(s) == 0 {
			continue
		}
		if _, exists := VALID_EXTENSIONS[s]; exists {
			exts = append(exts, s)
		}
	}
	return exts
}

func getExtensionRegex(extensions []string) string {
	return `.+\.(` + strings.Join(extensions, "|") + `)$`
}

func getFileListForPath(searchPath string, extensions []string) (fileList []string) {
	err := filepath.Walk(searchPath, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			logs.E("Process path %q error %v", currPath, err)
			return nil
		}
		if info.IsDir() {
			logs.V("skipping a dir without errors: %+v", info.Name())
			return nil
		}
		if FileExtensionMatchReg(currPath, extensions) {
			fileList = append(fileList, currPath)
		}
		logs.V("add file: %q\n", currPath)
		return nil
	})
	if err != nil {
		logs.E("Walking path %q error %v", searchPath, err)
		return
	}
	return
}

func getFileList(searchPath string, extensions []string) (fileList []string) {
	newFileList := getFileListForPath(searchPath, extensions)
	for _, file := range newFileList {
		file, _ = filepath.Rel(searchPath, file)
		fileList = append(fileList, file)
	}

	return
}

func getCheckIdNotIgnore(flagStr string) (checkerIds []int) {
	ignoreCheckerIds := getCheckerId(flagStr)
	for _, rule := range db.GetRules() {
		if ContainsInt(ignoreCheckerIds, rule.Id) {
			continue
		}
		checkerIds = append(checkerIds, rule.Id)
	}
	return
}
