package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("b   a    "))
}

func lengthOfLastWord(s string) int {
	//s = strings.TrimRight(s, " ")
	//return len(s) - strings.LastIndex(s, " ") - 1
	for s != "" && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	idx := -1
	for i := len(s); i > 0; i-- {
		if s[i-1] == ' ' {
			idx = i - 1
			break
		}
	}
	return len(s) - idx - 1
}
