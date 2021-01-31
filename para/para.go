package para

import (
	"flag"
	"github.com/tianshi25/additionalCheck/db"
	"github.com/tianshi25/additionalCheck/logs"
	"github.com/tianshi25/additionalCheck/tool"
	. "github.com/tianshi25/additionalCheck/filter"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func GetPara() (fileList []string, checkerIds []int, getInfoCheckerId int) {
	searchPath := flag.String("path", ".", "paths to check\ndefault:.")
	ignoreStr := flag.String("no", "", "ignore checker id:\nexample: 1,2\n"+db.GetBriefs())
	onlyStr := flag.String("only", "", "checker id\nexample: 1,2\n"+db.GetBriefs())
	extensionStr := flag.String("ext", "c,cpp,h,hpp,java,go", "file extension to check\ndefault: c,cpp,h,hpp,java,go")
	logLevel := flag.String("log", "Error", "log level\nvalue: Error Warn Info Verbose\ndefault:Error")
	getInfoCheckerIdAddr := flag.Int("info", 0, "get checker id info\n"+db.GetBriefs())
	flag.Parse()

	if len(*ignoreStr) != 0 &&  len(*onlyStr) != 0 {
		logs.E("ignore and only are conflict para")
		return
	}

	getInfoCheckerId = *getInfoCheckerIdAddr
	extensions := getExtensionList(*extensionStr)
	fileList = getFileList(*searchPath, extensions)
	if len(*ignoreStr) != 0 {
		checkerIds = getCheckIdNotIgnore(*ignoreStr)
	} else {
		checkerIds =getCheckerId(*onlyStr)
	}

	SetFilter(*searchPath)
	setLogLevel(*logLevel)
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
	regex := getExtensionRegex(extensions)
	err := filepath.Walk(searchPath, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			logs.E("Process path %q error %v", currPath, err)
			return nil
		}
		if info.IsDir() {
			logs.V("skipping a dir without errors: %+v", info.Name())
			return nil
		}
		if matched, err := regexp.MatchString(regex, filepath.Base(currPath)); err != nil {
			logs.E("error occur when match string")
			return err
		} else if matched {
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
		if !tool.ContainsString(fileList, file) {
			fileList = append(fileList, file)
		}
	}

	logs.I("files to check:" + strings.Join(fileList, "\n"))
	return
}

func getCheckIdNotIgnore(flagStr string) (checkerIds []int) {
	ignoreCheckerIds := getCheckerId(flagStr)
	for _, rule := range db.GetRules() {
		if tool.ContainsInt(ignoreCheckerIds, rule.Id) {
			continue
		}
		checkerIds = append(checkerIds, rule.Id)
	}
	return
}