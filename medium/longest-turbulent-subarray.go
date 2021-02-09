package main

import "fmt"

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 9}))
}

func maxTurbulenceSize(arr []int) int {
	/*
	   湍流子数组: 每隔一个数, 两个相邻数的大小关系都是反转的=>差值,都是正负相隔.

	   返回最长湍流子数组的长度=>最长湍流差值子数组长度+1
	*/
	if len(arr) <= 1 {
		return len(arr)
	}
	var res = make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		res[i] = arr[i+1] - arr[i]
	}
	var l = 0
	var max = l
	for i := 0; i < len(res); i++ {
		if i == 0 || !checkDiff(res[i], res[i-1]) {
			//i==0是特殊情况, 都是从该间隔重新计数的逻辑
			if res[i] != 0 {
				l = 1
			} else {
				l = 0
			}
		} else {
			l++
		}
		if l > max {
			max = l
		}
	}
	return max + 1
}

// 是否一正一负
func checkDiff(a, b int) bool {
	if a > 0 {
		return b < 0
	} else if a < 0 {
		return b > 0
	} else {
		return false
	}
}
