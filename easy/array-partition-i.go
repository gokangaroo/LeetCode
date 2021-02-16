package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 6, 7}))
}

func arrayPairSum(nums []int) (ans int) {
	/*
	   需要保证min(a,b)最大, 实际就是保证max(a,b)最小, 所以a和b在排序的nums中, 最接近就是最优解

	   直接排序后两两一组, 每一组都是最接近的数字, 大家都不追求最优的, 而是最合适的(对整体最优), 结果反而最优.

	   数学思维:
	       如果a<b<c<d, 那么a+c>a+b, 也就是隔着取(a,c)(b,d), 不如(a,b)(c,d), 画坐标轴, 也可以很容易发现, 前者有覆盖的差值部分, 是不合理的.
	*/
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return
}
