package main

import (
	"crawler/src/engine"
	"crawler/src/persist"
	"crawler/src/scheduler"
	"crawler/src/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		Count:     10,
		ItemChan: persist.ItemPersist(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
