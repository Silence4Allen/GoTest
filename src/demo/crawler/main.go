package main

import (
	"GoTest/src/demo/crawler/engine"
	"GoTest/src/demo/crawler/scheduler"
	"GoTest/src/demo/crawler/zhenai/parser"
	"GoTest/src/demo/crawler/persist"
)

var url = "http://www.zhenai.com/zhenghun"

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
