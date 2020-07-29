package myloger

import (
	"fmt"
	"time"
)

// 往终端写日志的相关内容

// Logger日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog的构造函数
func NewLog(levelStr string) ConsoleLogger {
	level := parseLevel(levelStr)
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(timeLayout)
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Waring(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
