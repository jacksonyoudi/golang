package parser

import (
	"golang/spider/engine"
	"regexp"
)

var cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


// 下一页 解析
func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)

	citys := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range citys {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]), ParserCity})
	}
	return result
}
