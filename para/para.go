package para

import (
    "tianshi25.github.com/additionalCheck/db"
    "tianshi25.github.com/additionalCheck/logs"
    "flag"
    "strings"
    "strconv"
    "path/filepath"
    "os"
    "regexp"
)

func getPara() (fileList []string, ignoreCheckerId []int) {
    path := flag.String("path", ".", "path to file or folder to check\n default:.")
    ignoreStr := flag.String("ignore", "", "ignore check:\n example: 1,2\n" + db.GetBriefs())
    extensionStr := flag.String("ext", "c,cpp,h,hpp,java,go", "file extension to check\ndefault: c,cpp,h,hpp,java,go")
    logLevel := flag.String("log", "Error", "log level\nvalue: Error Warn Info Verbose\ndefault:Error")

    extensions := getExtensionList(*extensionStr)
    fileList = getFileList(*path, extensions)
    ignoreCheckerId = getIgnoreCheckerId(*ignoreStr)
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
    for _, s := range(strings.Split(flagStr, ",")) {
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

var VALID_EXTENSIONS = map[string]int{"c":1, "h":1, "cpp":1, "hpp":1, "java":1}

func getExtensionList(extensionStr string) (exts []string) {
    for _, s := range(strings.Split(extensionStr, ",")) {
        if _, exists := VALID_EXTENSIONS[s]; exists {
            exts = append(exts, s)
        }
    }
    return exts
}

func getExtensionRegex(extensions []string) string {
    return `.*\.(` + strings.Join(extensions, "|") + `)$`
}

func getFileList(path string, extensions []string) (fileList []string) {
    regex := getExtensionRegex(extensions)
    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            logs.E("Process path %q error %v", path, err)
            return nil
        }
        if info.IsDir() {
            logs.V("skipping a dir without errors: %+v", info.Name())
            return nil
        }
        if matched, err := regexp.MatchString(regex, filepath.Base(path)); err != nil {
            return err
        } else if matched {
            fileList = append(fileList, path)
        }
        logs.V("add file: %q\n", path)
        fileList = append(fileList, path)
        return nil
    })
    if err != nil {
        logs.E("Walking path %q error %v", path, err)
        return
    }
    logs.I("files to check:" + strings.Join(fileList, "\n"))
    return
}