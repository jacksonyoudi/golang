package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type DebugInfo struct {
	Level string `json:"level,omitempty"` // Level解析level，忽略空值
	Msg string `json:"message"` // Msg解析为message
	Author string `json:"author"`
}


type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main()  {
	info := DebugInfo{
		"Info",
		"error",
		"helloWord"}

	str, _ := json.Marshal(info)
	fmt.Println(string(str))

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}

/*

可以选择的控制字段有三种：

-：不要解析这个字段
omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string
FieldName：当解析 json 的时候，使用这个名字

 */