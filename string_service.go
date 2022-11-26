package gokitstringservice

import "strings"

//StringService provides operations on strings
type StringService interface {
	//Uppercase return input sting as UPPERCASE
	Uppercase(string) (string, error)
	//Count length of input string
	Count(string) (int, error)
}

type stringService struct{}

func (stringService) Uppercase(input string) (output string, e error) {
	return strings.ToUpper(input), nil
}

func (stringService) Count(input string) (cnt int, e error) {
	return len(input), nil
}