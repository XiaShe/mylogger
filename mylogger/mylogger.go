package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Loglevel uint16

// 日志级别 高 ---> 低
const (
	UNKNOWN Loglevel = iota  //
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
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

func getlogString(lv Loglevel)string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	// 默认返回值
	return "DEBUG"
}


// 执行函数 行号、函数命、文件名判定
// 传入参数skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
func getInfo(n int) (funcName, fileName string, lineNo int ) {
	pc,file,lineNo,ok := runtime.Caller(n)
	// 如果不能成功执行runtime.Caller()方法：
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	// 获取函数名
	funcName = runtime.FuncForPC(pc).Name()
	// 获取文件名
	fileName = path.Base(file)
	// 将字符串按照.分割，取[1]字符
	funcName = strings.Split(funcName, ".")[1]
	return
}