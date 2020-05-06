package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	t := 0
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

// 先使用n^2解法看看
// 存在i,j满足条件
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ { //需要考虑k>len的场景
		for j := max(i-k, 0); j < i; j++ {
			if !(nums[j]-nums[i] < - 1*t || nums[j]-nums[i] > t) {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 利用排序特性=>可以优化成二叉搜索树版本, 或者改成二叉搜索删除以及插入
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	tmp := make(sortInts, 0)
	for i := 0; i < len(nums); i++ { //不能是len(nums)-k, 因为可能k更大
		if tmp.Len() == k+1 {
			//应该删掉开头那个数
			tmp.delXInSlice(nums[i-k-1])
		}
		if tmp.addXAndCompare(nums[i], t) {
			return true
		}
	}
	return false
}

// 利用排序特性, 比较最大最小值即可
// 可以利用二叉搜索树来实现删除和增加,以及比较
type sortInts []int

func (s *sortInts) Len() int {
	return len(*s)
}

// 删除某个值
func (s *sortInts) delXInSlice(x int) {
	left := 0
	right := s.Len() - 1
	for left <= right {
		mid := (left + right) / 2
		if (*s)[mid] == x { //进行剔除
			if mid == len(*s)-1 {
				*s = (*s)[:mid]
			} else {
				*s = append((*s)[:mid], (*s)[mid+1:]...)
			}
			return
		}
		if (*s)[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

// 添加某个值并判断是否满足条件
func (s *sortInts) addXAndCompare(x, t int) bool {
	left := 0
	right := s.Len() - 1
	if s.Len() == 0 {
		*s = append(*s, x)
		return false
	}
	if x >= (*s)[right] {
		*s = append(*s, x)
		return x-(*s)[right] <= t
	}
	if x <= (*s)[left] {
		*s = append([]int{x}, *s...)
		return (*s)[1]-x <= t
	}
	for left <= right {
		mid := (left + right) / 2
		if (*s)[left] <= x && (*s)[right] >= x {
			if x-(*s)[left] <= t || (*s)[right]-x <= t {
				return true
			} else {
				*s = append(append((*s)[:left+1], x), (*s)[right:]...)
			}
		}
		if (*s)[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
