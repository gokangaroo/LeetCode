package main

import "fmt"

func main() {
	//fmt.Println(myPow(0.00001, 2147483647))
	fmt.Println(myPow(2, -5))
	fmt.Println(myPow1(2, -5))
}

// 递归+二分法
func myPow(x float64, n int) float64 {
	var negative = n < 0
	var res = 1.0
	if n == 0 {
		return res
	}
	if n == 1 { //base case
		return x
	}
	if n%2 == 0 {
		if !negative { //不能简单两个myPow,而是将结果取出,相乘, 只需要计算一条边就行
			res = myPow(x, n/2)
			return res * res
		}
		res = myPow(x, -n/2)
		return 1 / res * res
	}
	if !negative {
		res = myPow(x, n/2)
		return res * res * x
	}
	res = myPow(x, -n/2)
	return 1 / (res * res * x)
}

// 迭代+二分法
func myPow1(x float64, n int) float64 {
	var negative = n < 0
	if n == 0 {
		return 1.0
	}
	var res = x
	if negative {
		n *= -1
	}
	var i = 1
	for ; i < n; i *= 2 {
		res *= res
	}
	//正向还是需要一个小递归
	res *= myPow(x, n-i)
	//反向方法 TODO 没什么必要, 难写.
	if negative {
		return 1 / res
	}
	return res
}
