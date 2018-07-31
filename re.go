package main

/*
regexp

 */

func main()  {
	//isok ,_ := regexp.Match("[a-zA-z]{3}", []byte("abc"))
	//isok ,_ := regexp.MatchString("[a-zA-z]{3}", "b1d")
	//fmt.Println(isok)

	reg := regexp.MustCompile("1233333abc")
	isok := reg.FindAllString("abc", 1)
	fmt.Println(isok)
}