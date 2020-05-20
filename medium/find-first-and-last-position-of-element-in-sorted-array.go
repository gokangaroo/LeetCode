package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 4}, 4))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)
	var mid int
	for left+1 < right { //模板3
		//left和right都是我们需要找的idx
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			//left在left->mid上
			//right在mid->right
			var midl, midr int
			var lmax, rmin = mid, mid
			for left < lmax { // 内部使用模板二,找到left==target
				midl = (left + lmax) / 2
				if nums[midl] < target {
					left = midl + 1
				} else {
					lmax = midl
				}
			}
			for rmin < right { //而right是大于的第一个数
				midr = (rmin + right) / 2
				if nums[midr] > target {
					right = midr
				} else {
					rmin = midr + 1
				}
			}
			break
		}
	}
	if nums[left] != target && nums[right-1] != target {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}
