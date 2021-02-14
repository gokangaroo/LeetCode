package main

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
}

func minSwapsCouples(row []int) int {
	/*
	   不需要想复杂(说是困难题, 想通下面两句话, 就变简单题目了)

	   每次交换, 至少可以成就1-2对情侣, 从贪心角度倾向2对, 但是1对, 还是2对, 无论我们怎么交换, 实际都不会改变(后面论证)
	   我们先保证每次交换都成全1对情侣, 也就是有目标的交换, 这样之后, 实际我们也能保证交换成全2对的两对座位, 其实并不会与其他对产生错误交换(因为没人需要你们, 反向保证了最优解)
	*/
	n, res := len(row), 0
	for i := 0; i < n; i += 2 { // i是座位编号
		x := row[i]
		if row[i+1] == getPair(x) {
			continue
		}
		res++ // 该对座位需要交换, for循环进行交换.
		for j := i + 1; j < n; j++ {
			if row[j] == getPair(x) {
				row[i+1], row[j] = row[j], row[i+1]
				break
			}
		}
	}
	return res
}

// 更简单的写法是x^1, 但是可读性差点.
func getPair(a int) int {
	if a%2 == 0 {
		return a + 1
	}
	return a - 1
}
