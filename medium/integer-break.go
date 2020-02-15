package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	fmt.Print(integerBreak(n))
}

// 遇到3和2才会拆分  其余根据2个数来做乘积
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	x := n % 3
	if x == 0 {
		return int(math.Pow(3, float64(n/3)))
	} else if x == 1 {
		return 4 * int(math.Pow(3, float64((n-4)/3)))
	} else {
		return 2 * int(math.Pow(3, float64((n-2)/3)))
	}
}
