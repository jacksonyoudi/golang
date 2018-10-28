package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city.html")

	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCitys := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}


	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d" +
			" reqeusts; but had %d", resultSize, len(result.Requests))
	}

	for i ,Url := range expectedUrls {
		if result.Requests[i].Url != Url {
			t.Errorf("Error")
		}
	}


	for i ,city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("Error")
		}
	}

}
