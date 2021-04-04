package main

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
}

func numRabbits(answers []int) int {
	/*
	   返回兔子的最少数量

	   1. 两只说出其他相同颜色数量一样的兔子, 可以理解为同颜色.
	   2. 先贪心的如此认为, 如果超过, 再分一个出来
	*/
	count := 0
	r := make(map[int]int) // answer => count(同answer的数量)
	for i := 0; i < len(answers); i++ {
		if c, ok := r[answers[i]]; ok {
			if c < answers[i]+1 { // 可以是同颜色的数量是answers[i]+1
				r[answers[i]]++
			} else {
				for r[answers[i]] >= answers[i]%1000+1 { // 当新槽内的数量还是超过时, 继续开辟新槽(设定新颜色).
					answers[i] += 1000 // gap
				}
				r[answers[i]]++
			}
		} else {
			r[answers[i]]++
		}
	}
	for k, _ := range r {
		for k >= 1000 {
			k -= 1000
		}
		count += k + 1 // 实际数量, 是answers[i]+1, k+1表示同颜色的数量
	}
	return count
}
