package main

import "fmt"

var pick = 6

func main() {
	var n = 10
	fmt.Println(guessNumber(n))
}

func guessNumber(n int) int {
	left, right := 1, n
	var mid int
	var res int
	for left <= right {
		mid = (left + right) / 2
		res = guess(mid)
		if res == 0 {
			break
		}
		if res == -1 {
			right = mid - 1
		}
		if res == 1 {
			left = mid + 1
		}
	}
	return mid
}

func guess(x int) int {
	if x < pick {
		return 1
	}
	if x > pick {
		return -1
	}
	return 0
}
