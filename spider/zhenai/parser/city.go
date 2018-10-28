package parser

import (
	"golang/spider/engine"
	"regexp"
)

const cityRe = `<a href="(http://[0-9a-z]+.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

func ParserCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)

	users := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range users {
		name := string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]), func(c []byte) engine.ParserResult {
				return ParseProfile(c, name)
			}})
	}
	return result
}
