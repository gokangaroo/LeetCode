package main

import "fmt"

func main() {
	var s = "PAYPALISHIRING"
	fmt.Println(convert(s, 7))
}

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows == 1 || len(s) <= numRows {
		return s
	}
	var s2 = ""
	res := make([]string, numRows)
	var x = 0
	var up = false //也就是down
	for i := 0; i < len(s); i++ {
		res[x] += string(s[i])
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
	for i := 0; i < numRows; i++ {
		s2 += res[i]
	}
	return s2
}
