package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		nums[(nums[i]-1)%n] += n
	}

	res := make([]int, 0, n)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res

}
