package main

import (
	"container/heap"
	"fmt"
	"sort"
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

// 使用小顶堆来实现, 小顶堆的Len保持为k
func findKthLargest2(nums []int, k int) int {
	n := IntHeap(nums[0:k])
	heap.Init(&n)
	var min = 0
	for index := k; index < len(nums); index++ {
		if min == 0 {
			min = heap.Pop(&n).(int)
		}
		if min < nums[index] { // 比堆顶大再推入(最大的那k个数形成的小顶堆),否则不推
			heap.Push(&n, nums[index])
			min = 0
		}
	}
	if min == 0 {
		return heap.Pop(&n).(int)
	}
	return min
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// <改成>就是大顶堆了
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 直接使用标准库的排序方法
func findKthLargest3(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
