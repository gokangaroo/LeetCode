package main

import (
	"fmt"
	"math"
)

func main() {
	i := -123
	fmt.Println(reverse(i))
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if (rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7)) {
			return 0
		}
		//if (rev > 0x7fffffff/10 || (rev == 0x7fffffff/10 && pop > 7)) {
		//	return 0
		//}
		// go没有int32溢出? 不然应该是0x80000000
		//if (rev < -0x7fffffff/10 || (rev == -0x7fffffff/10 && pop < -8)) {
		//	return 0
		//}
		if (rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8)) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
