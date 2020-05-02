package mylogger

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