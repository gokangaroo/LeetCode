package main

import "fmt"

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}

//暴力解法  会超时
func nthUglyNumber1(n int) int {
	//从2开始计算
	if n == 1 {
		return 1
	}
	i, j := 2, 1
	for {
		tmp := i
		for {
			for ; tmp%2 == 0; {
				tmp = tmp / 2
			}
			for ; tmp%3 == 0; {
				tmp = tmp / 3
			}
			for ; tmp%5 == 0; {
				tmp = tmp / 5
			}
			if tmp == 1 {
				j++
			}
			break
		}
		if j == n {
			return i
		}
		i++
	}
}

//三指针坐标法(2,3,5坐标点)
func nthUglyNumber(n int) int {
	nums := []int{1}
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := getMin(nums[i2]*2, getMin(nums[i3]*3, nums[i5]*5))
		if min == nums[i2]*2 {
			i2++
		}
		if min == nums[i3]*3 {
			i3++
		}
		if min == nums[i5]*5 {
			i5++
		}
		nums = append(nums,min)
	}
	return nums[n-1]
}

func getMin(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
