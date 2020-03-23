package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	res := permute(arr)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	for i, v := range nums {
		sub := make([]int, len(nums))
		copy(sub, nums)

		sub = append(sub[:i], sub[i+1:]...)
		// 递归,拼接出最终二维数组, range拼接
		for _, sr := range permute(sub) {
			rr := []int{v}
			rr = append(rr, sr...)
			res = append(res, rr)
		}
	}

	return res
}
