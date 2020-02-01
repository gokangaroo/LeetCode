package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "PAYPALISHIRING"
	fmt.Println(convert(s, 7))
}

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows == 1 || len(s) <= numRows {
		return s
	}
	res := make([]strings.Builder, numRows)
	var x = 0
	var up = false //也就是down
	for i := 0; i < len(s); i++ {
		res[x].WriteByte(s[i])
		if x < numRows-1 && !up {
			x++
			continue
		}
		if x == numRows-1 {
			x--
			up = true
			continue
		}
		if x > 0 && up {
			x--
			continue
		}
		if x == 0 {
			x++
			up = false
			continue
		}
	}
	var s2 strings.Builder
	for i := 0; i < numRows; i++ {
		s2.WriteString(res[i].String())
	}
	return s2.String()
}
