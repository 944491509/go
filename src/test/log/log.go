package main

import (
	"myloger"
	"time"
)

func main() {
	//log := myloger.NewLog("INFO")
	day := time.Now().Format("2006-01-02")
	log := myloger.NewFileLogger("INFO", "./Runtime/", day, 10*1024*1024)
	for {
		log.Debug("我是Debug日志")
		log.Info("我是Info日志")
		log.Waring("我是Waring日志")
		id := 130347
		name := "刘洋"
		log.Error("我是Error日志,id:%d, name:%s", id, name)
		time.Sleep(time.Second * 3)
	}

}
