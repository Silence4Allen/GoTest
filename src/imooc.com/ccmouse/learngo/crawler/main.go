package main

import (
	"GoTest/src/imooc.com/ccmouse/learngo/crawler/config"
	"GoTest/src/imooc.com/ccmouse/learngo/crawler/engine"
	"GoTest/src/imooc.com/ccmouse/learngo/crawler/persist"
	"GoTest/src/imooc.com/ccmouse/learngo/crawler/scheduler"
	"GoTest/src/imooc.com/ccmouse/learngo/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
