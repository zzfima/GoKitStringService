package gokitstringservice

type uppercaseRequest struct {
	Input string `json:"input"`
}

type uppercaseResponse struct {
	Output string `json:"output"`
	E      error  `json:"e"`
}

type countRequest struct {
	Input string `json:"input"`
}

type countResponse struct {
	Cnt int `json:"cnt"`
}
