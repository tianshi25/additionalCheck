package para

import (
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"tianshi25.github.com/additionalCheck/db"
	"tianshi25.github.com/additionalCheck/logs"
	"tianshi25.github.com/additionalCheck/tool"
)

func GetPara() (fileList []string, ignoreCheckerIds []int, getInfoCheckerId int) {
	searchPaths := flag.String("path", ".", "paths to check\ndefault:.\nexample:./1,./2")
	ignoreStr := flag.String("ignore", "", "ignore checker id:\nexample: 1,2\n"+db.GetBriefs())
	extensionStr := flag.String("ext", "c,cpp,h,hpp,java,go", "file extension to check\ndefault: c,cpp,h,hpp,java,go")
	logLevel := flag.String("log", "Error", "log level\nvalue: Error Warn Info Verbose\ndefault:Error")
	getInfoCheckerIdAddr := flag.Int("info", 0, "get checker id info\n"+db.GetBriefs())
	flag.Parse()

	getInfoCheckerId = *getInfoCheckerIdAddr
	extensions := getExtensionList(*extensionStr)
	fileList = getFileList(*searchPaths, extensions)
	ignoreCheckerIds = getIgnoreCheckerId(*ignoreStr)
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

func getIgnoreCheckerId(flagStr string) []int {
	ids := []int{}
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

func getFileList(searchPaths string, extensions []string) (fileList []string) {
	for _, path := range strings.Split(searchPaths, ",") {
		if len(path) == 0 {
			continue
		}
		new_file_list := getFileListForPath(path, extensions)
		for _, file := range new_file_list {
			if !tool.ContainsString(fileList, file) {
				fileList = append(fileList, file)
			}
		}

	}
	logs.I("files to check:" + strings.Join(fileList, "\n"))
	return
}
