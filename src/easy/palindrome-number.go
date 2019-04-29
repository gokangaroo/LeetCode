package main

import "fmt"

func main() {
	x := 1001
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}
	m := x
	y := 0
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return m == y
}
