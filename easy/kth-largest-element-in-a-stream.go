package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	arr := []int{4, 5, 8, 2}
	kth := Constructor(3, arr)
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}

// 维护一个k-1长度的最小堆即可
type KthLargest struct {
	m   *MinHeap
	kth int
}

func Constructor(k int, nums []int) KthLargest {
	var res = KthLargest{
		m:   &MinHeap{},
		kth: math.MinInt64,
	}
	if len(nums) >= k {
		*res.m = nums[:k]
		heap.Init(res.m)
		res.kth = heap.Pop(res.m).(int)
		for i := k; i < len(nums); i++ {
			res.Add(nums[i])
		}
	} else {
		*res.m = nums
		heap.Init(res.m)
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	if val > this.kth { //不是一定要堆, 完全可以append, 然后直接排序.
		heap.Push(this.m, val)
		this.kth = heap.Pop(this.m).(int)
	}
	return this.kth
}

// 小顶堆
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
