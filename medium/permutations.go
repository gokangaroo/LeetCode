package main

import "fmt"

func main()  {
	arr := []int{1,2,3}
	res := permute(arr)
	fmt.Println(res)
}

func permute(nums []int) [][]int{
	res := make([][]int,0)
	if len(nums) == 1{
		res = append(res,nums)
		return res
	}

	for i := range nums{
		sub := make([]int,len(nums))
		copy(sub,nums)

		sub = append(sub[:i],sub[i+1:]...)
		subR := permute(sub)
		for _,sr := range subR {
			rr := []int{nums[i]}
			rr = append(rr, sr...)
			res = append(res,rr)
		}
	}

	return res
}