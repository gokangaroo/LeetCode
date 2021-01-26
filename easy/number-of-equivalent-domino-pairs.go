package main

import (
	"fmt"
)

func main() {
	fmt.Println(numEquivDominoPairs([][]int{
		{1, 2}, {2, 1}, {3, 4}, {5, 6},
	}))
}

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	// 1-9的数字, 排序后组合成二位数, 然后计数
	cnt := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		// ans += cnt[v]
		cnt[v]++
	}
	// 计算对数比较直接, 但是上面的加法更妙
	for _,v:=range cnt{
		ans=ans+v*(v-1)/2
	}
	return
}
