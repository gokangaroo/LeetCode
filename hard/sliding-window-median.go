package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(medianSlidingWindow([]int{2147483647, 1, 2, 3, 4, 5, 6, 7, 2147483647}, 2))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	/*
	   每一次移动都是不同的切片
	   比较好的方法是, 维护两个堆
	   数量相等, 一个大顶堆是左侧, 一个小顶堆是右侧
	   保证大顶堆堆顶<=小顶堆堆顶, 可以全入一边, 然后pop一半到另一边

	   问题: 被区间去掉的数字如何Pop出来=>堆如何Pop指定Value=>可以换个思路, 进行修改后fix
	*/
	var (
		res []float64
		m   = make(map[int]*Item, k)
	)
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			res = append(res, float64(nums[i]))
		}
		return res
	}
	// 左右侧
	left := make(DESCQueue, k)
	right := make(ASCQueue, k/2)
	for i := 0; i < k; i++ {
		item := &Item{value: nums[i], index: i}
		left[i] = item
		m[i] = item
	}
	heap.Init(&left)
	// 数量相等
	for i := 0; i < k/2; i++ {
		item := heap.Pop(&left).(*Item)
		item.index = i
		right[i] = item
	}
	heap.Init(&right)
	// 求中位数
	res = append(res, getMid(left, right, k))
	for i := k; i < len(nums); i++ {
		// 替换m
		head := m[i-k]
		delete(m, i-k)
		m[i] = head
		tail := nums[i]
		// 更新内容
		if left[head.index] == head {
			left.Update(head, tail)
		} else {
			right.Update(head, tail)
		}
		exchange(left, right)
		// 再求中位数
		res = append(res, getMid(left, right, k))
	}
	return res
}

// 确认是否需要交换堆顶
func exchange(left DESCQueue, right ASCQueue) {
	lMax := heap.Pop(&left).(*Item)
	rMin := heap.Pop(&right).(*Item)
	if lMax.value > rMin.value {
		// exchange
		heap.Push(&left, rMin)
		heap.Push(&right, lMax)
	} else {
		// no change
		heap.Push(&left, lMax)
		heap.Push(&right, rMin)
	}
}

// 获取两个堆的中位数
func getMid(left DESCQueue, right ASCQueue, k int) float64 {
	mid := heap.Pop(&left).(*Item)
	defer heap.Push(&left, mid)
	if k%2 == 0 {
		mid2 := heap.Pop(&right).(*Item)
		defer heap.Push(&right, mid2)
		return float64(mid.value)/2 + float64(mid2.value)/2
	}
	return float64(mid.value)
}

type Item struct {
	value int
	index int
}

type ASCQueue []*Item  //小顶
type DESCQueue []*Item //大顶

func (aq ASCQueue) Len() int  { return len(aq) }
func (dq DESCQueue) Len() int { return len(dq) }

func (aq ASCQueue) Less(i, j int) bool {
	return aq[i].value < aq[j].value
}
func (dq DESCQueue) Less(i, j int) bool {
	return dq[i].value > dq[j].value
}

func (aq ASCQueue) Swap(i, j int) {
	aq[i], aq[j] = aq[j], aq[i]
	aq[i].index = i
	aq[j].index = j
}
func (dq DESCQueue) Swap(i, j int) {
	dq[i], dq[j] = dq[j], dq[i]
	dq[i].index = i
	dq[j].index = j
}

func (aq *ASCQueue) Push(x interface{}) {
	n := len(*aq)
	item := x.(*Item)
	item.index = n
	*aq = append(*aq, item)
}
func (dq *DESCQueue) Push(x interface{}) {
	n := len(*dq)
	item := x.(*Item)
	item.index = n
	*dq = append(*dq, item)
}

func (aq *ASCQueue) Pop() interface{} {
	old := *aq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*aq = old[0 : n-1]
	return item
}
func (dq *DESCQueue) Pop() interface{} {
	old := *dq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*dq = old[0 : n-1]
	return item
}

func (aq *ASCQueue) Update(item *Item, value int) {
	item.value = value
	heap.Fix(aq, item.index)
}
func (dq *DESCQueue) Update(item *Item, value int) {
	item.value = value
	heap.Fix(dq, item.index)
}
