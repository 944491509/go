package myloger

// 网文件里面写日志

type FileLogger struct {
	Level LogLevel
	filePath string  // 日志保存路径
	fileName string  // 日记文件名称
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger  {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return &FileLogger{
		Level: LogLevel,
		filePath: fp,
		fileName: fn
		maxFileSize: maxSize,
	}
}