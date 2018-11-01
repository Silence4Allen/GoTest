package main

import (
	"GoTest/src/demo/crawler/engine"
	"GoTest/src/demo/crawler/scheduler"
	"GoTest/src/demo/crawler/zhenai/parser"
	"GoTest/src/demo/crawler/persist"
)

var url = "http://www.zhenai.com/zhenghun"

func main() {
	//itemChan是需要展示数据的通道
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	//配置引擎（调度器、协程数量、数据传输的通道）
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	//引擎启动工作
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCity,
	})
}
