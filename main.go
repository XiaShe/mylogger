package main

import (
	"study1/mylogger"
)

// 声明全局接口变量
var log mylogger.Logger

func main() {
	// 终端日志实例，限定输出error以下级别的日志
	//log := mylogger.NewConsoloLogger("error")

	/*
		文件日志实例，限定输出info以下级别的日志，其中：
		levelStr：限定输出日志的级别
		fp：filepath 日志路径
		fn：filename 日志名称
		maxSize：每个日志文件的大小
	*/
	log = mylogger.NewFileLogger("info", "./", "xiashe.log", 10*1024*1024)

	// 下面 a n 表示写入日志的内容，任意写入参数数量与内容（空接口接受）
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		a := "我错了"
		n := "我真的错了"
		log.Error("这是一条Error日志，%s,%s", a, n)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
