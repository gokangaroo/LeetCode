package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))
}

func sumFourDivisors(nums []int) int {
	/*
	   找到恰好4个因数的数字, 并返回这些因数的和
	*/
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + fourDivisorsSum(nums[i])
	}
	return sum
}

func fourDivisorsSum(num int) int {
	count := 2
	sum := 1 + num
	for i := 2; i < mySqrt(num)+1; i++ {
		if num%i == 0 {
			if i*i == num || count >= 4 {
				return 0
			}
			count = count + 2
			sum = sum + i + num/i
		}
	}
	if count == 4 {
		return sum
	}
	return 0
}

func mySqrt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}
