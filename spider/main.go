package main

import (
	"golang/spider/engine"
	"golang/spider/persist"
	"golang/spider/scheduler"
	"golang/spider/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		&scheduler.SimpleScheduler{}, 100,
		persist.ItemSaver()}
	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun", parser.ParserCityList})
}
