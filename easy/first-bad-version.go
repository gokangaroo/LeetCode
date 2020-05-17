package main

import "fmt"

var aim = 1

func main() {
	fmt.Println(firstBadVersion(2))
}

func firstBadVersion(n int) int {
	left, right := 1, n
	var mid int
	for left < right { // aim是left==right
		mid = (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(x int) bool {
	return x >= aim
}
