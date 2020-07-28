package myloger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

const timeLayout = "2006-01-02 15:04:05"

func parseLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "waring":
		return WARNING
	case "error":
		return ERROR
	default:
		return UNKNOWN
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	}
	return "DEBUG"
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller failed , err : ")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
