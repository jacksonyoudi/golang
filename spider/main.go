package main

import (
	"golang/spider/engine"
	"golang/spider/scheduler"
	"golang/spider/zhenai/parser"
)


func main() {
	e := engine.ConcurrentEngine{&scheduler.SimpleScheduler{}, 100}
	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun", parser.ParserCityList})
}
