package parser

import (
	"fmt"
	"golang/spider/engine"
	"golang/spider/model"
	"regexp"
	"strconv"
)

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Endcation  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([0-9]+)CM</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)元</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var eduRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {

	profile := model.Profile{}
	profile.Name = name
	// 年龄
	age, err := strconv.Atoi(extracString(contents, ageRe))

	if err == nil {
		profile.Age = age
	}

	//
	height, err := strconv.Atoi(extracString(contents, heightRe))

	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extracString(contents, weightRe))

	if err == nil {
		profile.Weight = weight
	}

	income := extracString(contents, incomeRe)

	if err == nil {
		profile.Income = income
	}

	gender := extracString(contents, genderRe)

	if err == nil {
		profile.Gender = gender
	}

	marriage := extracString(contents, marriageRe)

	if err == nil {
		profile.Marriage = marriage
	}

	edu := extracString(contents, eduRe)

	if err == nil {
		profile.Endcation = edu
	}

	hokou := extracString(contents, hokouRe)

	if err == nil {
		profile.Hokou = hokou
	}

	house := extracString(contents, houseRe)

	if err == nil {
		profile.House = house
	}

	car := extracString(contents, carRe)

	if err == nil {
		profile.Car = car
	}

	xinzuo := extracString(contents, xinzuoRe)

	if err == nil {
		profile.Xinzuo = xinzuo
	}

	occupation := extracString(contents, occupationRe)

	if err == nil {
		profile.Occupation = occupation
	}

	fmt.Printf("%v", profile)
	result := engine.ParserResult{
		Items: engine.Item{
			Url: "",
			Id:  "", Type: "zhenai", Payload: profile},
	}
	return result
}

func extracString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
