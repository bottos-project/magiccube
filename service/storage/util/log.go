package util

import (
	"runtime"
	"fmt"
	"strings"
)

func FuncLog(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File: %s  Function: %s Line: %d", path(file), runtime.FuncForPC(function).Name(), line)
}

func path(filePath string) string {
	i := strings.LastIndex(filePath, "/")
	if i == -1 {
		return filePath
	} else {
		return filePath[i+1:]
	}
}