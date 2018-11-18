package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []Item
}




func NilParser([]byte) ParserResult {
	return ParserResult{}
}

type Item struct {
	Url string
	Id string
	Type string
	Payload interface{}
}
