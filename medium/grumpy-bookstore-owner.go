package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied(
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	/*
	   1.len(customers)分钟
	   2.第i分钟customers[i]顾客
	   3.grumpy[i]==1表示在i分钟生气=>顾客不满意
	   4.老板可以连续X分钟不生气=>最多满意数

	   也就是最多能够由不满意转满意的人, 假设原满意人数为ori, 最多转满意人数为tmax
	   固定X分钟=>滑动窗口
	*/
	tmax,ori:=0,0
	rangeT:=0

	for i:=0;i<X;i++{
		if grumpy[i]==0{
			ori += customers[i]
		}else{
			rangeT += customers[i]
		}
	}
	tmax = rangeT

	for i:=X;i<len(customers);i++{//假设X分钟,是[i-X+1,i]
		if grumpy[i]==0{
			ori += customers[i]
		}else{
			rangeT += customers[i]
		}

		if grumpy[i-X]==1{
			rangeT -= customers[i-X]
		}

		if rangeT>tmax{
			tmax=rangeT
		}
	}
	return ori+tmax
}
