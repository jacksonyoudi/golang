package parser

import (
	"golang/spider/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	content, err := fetcher.Fetch("http://album.zhenai.com/u/102930423")
	if err != nil {
		panic(err)
	}

	ParseProfile(content, "fz")

}
