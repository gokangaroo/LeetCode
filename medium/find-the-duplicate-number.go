package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	// 找不重复的可以使用异或计算
	// 1. 找重复的, 时间复杂度n方, 那就双指针就好了
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] == nums[j] {
	// 			return nums[i]
	// 		}
	// 	}
	// }
	// 2. 排序n log(n)
	// sort.Ints(nums)
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		return nums[i]
	// 	}
	// }
	// 3. 使用下标, 数字1-n, 下标0-n, 肯定有一个数字有两个数字作为指向它(当成链表)
	// 那么以链表来看, 重复的数字, 一定处于环形的开口.
	// 这样的解法, 跳跃性比较强, 无法很好使用cpu缓存, 但是时间复杂度O(n)
	// n,L,n-L: 2(n-L+b)=(n-L+b)+kL (快慢指针相遇后,b表示到环口后满指针继续走的距离)
	// n-L=kL-b: 快指针再走n-L步, 总共走n-L+kL也就是会到达环口.
	i, j := nums[0], nums[nums[0]]
	for i != j { //走了kL步
		i = nums[i]       //slow
		j = nums[nums[j]] //fast
	}
	i = 0
	for i != j { //快指针再走n-L步, 到达环口
		i = nums[i]
		j = nums[j]
	}
	return i
}
