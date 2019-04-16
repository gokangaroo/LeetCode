package main

import (
	"datastructure"
	"fmt"
)

func main() {
	str := "()"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	cs := []byte(s)
	st := make(datastructure.Stack, 0)
	for _, c := range cs {
		switch c {
		case '(', '{', '[':
			st.Push(c)
		case ']':
			pope, _ := st.Pop()
			if []byte("[")[0] == pope {
			} else {
				return false
			}
		case '}':
			pope, _ := st.Pop()
			if []byte("{")[0] == pope {
			} else {
				return false
			}
		case ')':
			pope, _ := st.Pop()
			if []byte("(")[0] == pope {
			} else {
				return false
			}
		}
	}
	if st.Len() == 0 {
		return true
	} else {
		return false
	}
}