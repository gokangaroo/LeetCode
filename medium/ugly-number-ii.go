package main

import "fmt"

func main() {
	n := 13
	fmt.Println(nthUglyNumber1(n))
}

//暴力解法  会超时
func nthUglyNumber1(n int) int {
	i, j := 2, 1
	//从2开始计算
	if n == 1 {
		return 1
	}
	for {
		tmp := i
		for {
			for ; tmp%2 == 0; {
				tmp = tmp / 2
			}
			for ; tmp%3 == 0; {
				tmp = tmp / 3
			}
			for ; tmp%5 == 0; {
				tmp = tmp / 5
			}
			if tmp == 1 {
				j++
			}
			break
		}
		if j == n {
			return i
		}
		i++
	}
}