package mylogger

import (
	"errors"
	"fmt"
	"time"
)

// 判断日志级别
func parseLogLevel(s string) (Loglevel, error)  {
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效日志级别")
		return UNKNOWN, err
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

