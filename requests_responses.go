package gokitstringservice

type uppercaseRequest struct {
	Input string `json:"input"`
}

type uppercaseResponse struct {
	Output string `json:"output"`
	Err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	Input string `json:"input"`
}

type countResponse struct {
	Cnt int `json:"cnt"`
}
