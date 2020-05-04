package main

import (
	"study1/mylogger"
	"time"
)

func main() {
	// 限定输出error以下级别的日志
	log := mylogger.Newlog("error")
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}