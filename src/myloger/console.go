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

func log(lv LogLevel, msg string) {
	now := time.Now().Format(timeLayout)
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Waring(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
