package nlog

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetTrace(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", "unknown", 0
	}

	funcName := runtime.FuncForPC(pc).Name()

	if idx := strings.LastIndex(funcName, "/"); idx != -1 {
		funcName = funcName[idx+1:]
	}

	return filepath.Base(file), funcName, line
}
