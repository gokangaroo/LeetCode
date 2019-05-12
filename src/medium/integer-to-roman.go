package main

import "fmt"

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	special := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	svalue := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	mymap := make(map[int]string, 13)
	for i := 0; i < len(special); i++ {
		mymap[svalue[i]] = special[i]
	}
	sb := ""
	y := num
	for i := len(special) - 1; i >= 0; i-- {
		num = y / svalue[i]
		y = y % svalue[i]
		for j := 0; j < num; j++ {
			sb += mymap[svalue[i]]
		}
		if y == 0 {
			return sb
		}
	}
	return sb
}
