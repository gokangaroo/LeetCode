package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 1, 5, 6, 4, 8, 10}
	k := 3
	res := findKthLargest(arr, k)
	fmt.Println(res)
}

//堆排序的场景  第几个也就是第几次建堆
func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		//一次堆排序
		for j := len(nums) - 1; j >= 0; j-- {
			headily(nums, j)
		}
		if i != k-1 {
			nums = nums[1:]
		}
	}
	return nums[0]
}

//堆排序(大堆)
func headily(nums []int, currentIndex int) {
	size := len(nums)

	max := currentIndex
	//左右节点位置
	left := currentIndex*2 + 1
	right := currentIndex*2 + 2

	if left < size {
		if nums[max] < nums[left] {
			max = left
		}
	}
	if right < size {
		if nums[max] < nums[right] {
			max = right
		}
	}
	//如果最大的不是根元素 那么久交换
	if max != currentIndex {
		tmp := nums[max]
		nums[max] = nums[currentIndex]
		nums[currentIndex] = tmp

		//继续比较
		headily(nums, max)
	}
}
