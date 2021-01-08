package main

import "fmt"

func main() {
	is := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(is, 3)
	fmt.Println(is)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	// 环状替代
	var cycle = k
	for i := 0; cycle != 0; i++ {
		tmp := nums[i]
		var start, next int
		for start = i; (start+k)%len(nums) != i; start = next {
			next = (start + k) % len(nums)
			if next < start {
				cycle--
			}
			tmp, nums[next] = nums[next], tmp
		}
		cycle--
		nums[i] = tmp
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	// 翻转整个->翻转[0:k]->翻转[k:]
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; j > i; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums)
	// 整段copy
	tmp := make([]int, k)
	copy(tmp[:], nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], tmp[:])
}

func rotate1(nums []int, k int) {
	// 从最后一个开始, 逐个往前替换, k表示替换k个
	k %= len(nums)
	for i := 0; i < k; i++ {
		rotateOnce(nums)
	}
}

func rotateOnce(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
}
