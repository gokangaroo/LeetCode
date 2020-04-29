package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := make([]int, 1)
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	//  ¯\_(ツ)_/¯
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针, 一个是将nums1的数全部挪到后面(前往后), 一个是从后往前
// 第二种从后往前不需要提交对nums1进行操作
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for i >= 0 || j >= 0 {
		if j == -1 { //nums2全部排over
			return
		}
		if i == -1 {
			copy(nums1, nums2[:j+1])
			return
		}
		if nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
			continue
		} else {
			nums1[i+j+1] = nums2[j]
			j--
			continue
		}
	}
}
