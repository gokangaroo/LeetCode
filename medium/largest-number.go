package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nums = []int{0,0,0}
	res := largestNumber(nums)
	fmt.Println(res)
}

func largestNumber(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	length := len(nums)
	var res strings.Builder
	for i := 1; i < length; i++ {
		for j := length - 1; j >= i; j-- {
			if ifSwap(nums[j-1], nums[j]) {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
		if res.String() == "0" && nums[i-1] == 0 {
			continue
		}
		res.WriteString(strconv.Itoa(nums[i-1]))
	}
	if res.String() != "0" || nums[length-1] != 0 {
		res.WriteString(strconv.Itoa(nums[length-1]))
	}


	return res.String()
}

//是否需要交换
func ifSwap(a, b int) bool {
	x := strconv.Itoa(a) + strconv.Itoa(b)
	y := strconv.Itoa(b) + strconv.Itoa(a)

	if strings.Compare(x, y) == -1 && x < y {
		return true
	}
	return false
}
