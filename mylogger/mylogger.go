package mylogger

import "errors"

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
