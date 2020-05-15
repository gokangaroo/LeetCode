package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	left, right := 0, x/2+1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if mid*mid == x {
			break
		}
		if mid*mid < x {
			left = mid + 1
		}
		if mid*mid > x {
			right = mid - 1
		}
	}
	if mid*mid > x {
		return mid - 1
	}
	return mid
}
