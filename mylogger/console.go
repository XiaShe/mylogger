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

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [Debug] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}


func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}


func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [Fatal] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

