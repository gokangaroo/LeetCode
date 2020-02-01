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
	res := make([][]string, numRows)
	for index := 0; index < numRows; index++ {
		res[index] = make([]string, len(s)/(numRows-1)+1)
	}
	var x, y = 0, 0
	var up = false //也就是down
	for i := 0; i < len(s); i++ {
		res[x][y] = string(s[i])
		if x < numRows-1 && !up {
			x++
			continue
		}
		if x == numRows-1 {
			x--
			y++
			up = true
			continue
		}
		if x > 0 && up {
			x--
			continue
		}
		if x == 0 {
			x++
			y++
			up = false
			continue
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < y+1; j++ {
			if res[i][j] != "" {
				s2 += res[i][j]
			}
		}
	}
	return s2
}
