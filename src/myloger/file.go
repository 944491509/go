package myloger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 网文件里面写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志保存路径
	fileName    string // 日记文件名称
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel := parseLevel(levelStr)
	//if err != nil {
	//	panic(err)
	//}

	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	// 按照文件路径和文件名称将文件打开
	err := fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed , err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+"_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed , err:%v\n", err)
		return err
	}

	// 日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(timeLayout)
		funcName, fileName, lineNo := getInfo(3)
		if lv < ERROR {
			fmt.Fprintf(f.fileObj, "[%s] [%s][%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		}
		if lv >= ERROR {
			// 如果记录的日志大于等于error级别,我还要在err日志文件中再记录一边
			fmt.Fprintf(f.errFileObj, "[%s] [%s][%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)

		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Waring(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
