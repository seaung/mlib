package mlib

import (
	"io"
	"log"
	"os"
)

// 设置日志格式
const (
	flag     = log.Ldate | log.Ltime | log.Lshortfile // [DEBUG]2020/12/01 11:33:07 file.go:43: message
	LInfo    = "[INFO]"
	LDebug   = "[DEBUG]"
	LWarning = "[WARNING]"
	LError   = "[ERROR]"
)

// 定义日志器
var (
	logFile        io.Writer
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errrorLogger   *log.Logger
	defaultLogFile = "/tmp/log/yongda.log"
)

// 初始化日志器
func init() {
	var err error

	logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		defaultLogFile = "./yongda.log"
		logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("create log file error %+v\n", err)
		}

	}

	debugLogger = log.New(logFile, LDebug, flag)
	infoLogger = log.New(logFile, LInfo, flag)
	warningLogger = log.New(logFile, LWarning, flag)
	errrorLogger = log.New(logFile, LError, flag)
}

// 更改日志文件的输出路径
func SetLogOutputPath(path string) {
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("create log file error : %+v\n", err)
	}

	debugLogger.SetOutput(logFile)
	infoLogger.SetOutput(logFile)
	warningLogger.SetOutput(logFile)
	errrorLogger.SetOutput(logFile)
}

// 记录deug相关日志信息
func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

// 记录重要相关的日志信息
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// 记录警告相关的日志信息
func Warningf(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

// 记录错误相关的日志信息
func Errorf(format string, v ...interface{}) {
	errrorLogger.Printf(format, v...)
}
