package main

import "fmt"

func main() {
	xx := []int{1, 2, 1}
	fmt.Println(singleNumber(xx))
}

func singleNumber(nums []int) int {
	// 如何让自己跟自己计算永远抵消?
	// 异或运算,不一样才为1,最后都会抵消
	// 利用此特性也可以用来进行数值交换: a=a^b, b=a^b, a=a^b// a=a^b^b^(a^b). b=a^b^b
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
