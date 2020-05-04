package mylogger

import (
	"fmt"
	"time"
)

// Logger 日志结构体
type Logger struct {
	Level Loglevel
}

// Newlog 构造函数
func Newlog(levelStr string) Logger {
	// 判断日志级别，错误输入则报错
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return Logger{
		Level:level,
	}
}

// 输入日志界别判断
func (l Logger) enable(loglevel Loglevel) bool {
	return loglevel >= l.Level
}

//
func log(lv Loglevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
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

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}


func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}


func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}

