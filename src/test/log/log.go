package main

import (
	"myloger"
	"time"
)

func main() {
	log := myloger.NewLog("INFO")
	for {
		log.Debug("我是Debug日志")
		log.Info("我是Info日志")
		log.Waring("我是Waring日志")
		log.Error("我是Error日志")
		time.Sleep(time.Second * 3)
	}

}
