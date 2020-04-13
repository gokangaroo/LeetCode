package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(recursive)
}

var recursive = make(map[int]int)

func climbStairs(n int) int {
	if res, ok := recursive[n]; ok {
		return res
	}
	if n == 0 || n == 1 || n == 2 { //base case
		recursive[n] = n
	} else {
		recursive[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	// 直接递归会超时, 需要记忆
	// return climbStairs(n-1)+climbStairs(n-2)
	return recursive[n]
}
