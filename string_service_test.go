package gokitstringservice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	var s1 stringService
	c, _ := cnt(s1, "abc")
	require.Equal(t, 3, c)
}

func cnt(strsvc StringService, input string) (int, error) {
	return strsvc.Count(input)
}

func ExampleCount() {
	var s1 stringService
	sentence := "Mama hello"
	c, _ := s1.Count("Mama hello")
	fmt.Println("Sentence: " + sentence + " contains " + fmt.Sprint(c) + " words")

	//Output:
	//Sentence: Mama hello contains 10 words
}
