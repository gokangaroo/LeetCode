package main

import "fmt"

func main() {
	//fmt.Println(findClosestElements([]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4))
	//fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}

func findClosestElements(nums []int, k int, x int) []int {
	// 最靠近x的k个数(这个最靠近, 不是idx最近, 而是差值最近..)
	// 假设全部减去x, 改成最靠近0的数字
	if nums[0] >= x {
		return nums[:k]
	}
	if nums[len(nums)-1] <= x {
		return nums[len(nums)-k:]
	}
	//找到刚好>=x的第一个数,二分法模板二
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// 窗口初始位置
	if k%2 == 0 {
		left = left - k/2
		right = right + k/2
	} else {
		left = left - (k-1)/2 - 1
		right = right + (k-1)/2
	}
	// 窗口在范围外的场景重新定义窗口
	if left <= 0 {
		left = 1 //预留一个
		right = k + 1
	}
	if right >= len(nums) {
		left = len(nums) - k - 1
		right = len(nums) - 1 //也预留一个
	}
	// 开始平移窗口
	for left-1 >= 0 && right <= len(nums)-1 {
		if x-nums[left-1] <= nums[right-1]-x { //使用left-1代替right-1, 加个等于, 小数优先
			left--
			right--
		} else if nums[right]-x < x-nums[left] {
			left++
			right++
		} else {
			break
		}
	}
	return nums[left:right]
}
