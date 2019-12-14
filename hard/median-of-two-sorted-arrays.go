package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

//题目要求, 时间复杂度O(log(m + n)), 这边是有序数组, 那还好.
func findMedianSortedArrays(A []int, B []int) float64 {
	m := len(A)
	n := len(B)
	if m > n { // 确保m<=n. 直接交换两个数组
		temp := A
		A = B
		B = temp
		tmp := m
		m = n
		n = tmp
	}
	// 主要思想, 在A中找到一个i划分, 在B中找到一个j划分
	// A左侧和B左侧是中位数往左, A右侧和B右侧是中位数往右.
	// 需要兼容两者和为奇偶数的情况: 奇数, Max(A[i],B[j]), 偶数A[i]+B[j]/2
	// A[i-1]<B[j]并且B[j-1]<A[i]
	// i+j=m+n-i-j
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && B[j-1] > A[i] {
			iMin = i + 1 // i太小, A[i]不够大
		} else if i > iMin && A[i-1] > B[j] {
			iMax = i - 1 // j太大, B[j]不够小
		} else { // 找到了临界点, 进行情况分析.
			maxLeft := 0
			if i == 0 { // 说明A全在右边, 中位树在B
				maxLeft = B[j-1]
			} else if j == 0 { // 反之中位数在A
				maxLeft = A[i-1]
			} else {
				maxLeft = If(A[i-1] > B[j-1], A[i-1], B[j-1]).(int) // 混杂情况
			}
			if (m+n)%2 == 1 { // 如果是奇数, 直接返回即可
				return float64(maxLeft)
			}
			minRight := 0
			if i == m { // A全在左边, 中位数在B
				minRight = B[j]
			} else if j == n { // 反之中位数在A
				minRight = A[i]
			} else {
				minRight = If(B[j] < A[i], B[j], A[i]).(int)
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

//模拟一个三元运算符
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
