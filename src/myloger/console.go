package myloger

import (
	"fmt"
	"time"
)

// 往终端写日志的相关内容

// Logger日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog的构造函数
func NewLog(levelStr string) Logger {
	level := parseLevel(levelStr)
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format(timeLayout)
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)
	}
}

func (l Logger) Waring(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)
	}
}
