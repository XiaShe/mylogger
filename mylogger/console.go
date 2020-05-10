package mylogger

import (
	"fmt"
	"time"
)

// 终端输出 日志结构体
type ConsoleLogger struct {
	Level Loglevel // 日志级别
}

// NewConsoloLogger 构造函数
func NewConsoloLogger(levelStr string) ConsoleLogger {
	// 判断日志级别，错误输入则报错
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	// 没有错误则返回日志级别
	return ConsoleLogger{
		Level: level,
	}
}

// 根据日志级别（loglevel），判断是否需要记录该日志
func (c ConsoleLogger) enable(loglevel Loglevel) bool {
	return loglevel >= c.Level
}

// 在终端上打印日志信息
func (c ConsoleLogger) log(lv Loglevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
	}
}

// 下面是将不同级别日志打印到终端的方法

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)

}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
