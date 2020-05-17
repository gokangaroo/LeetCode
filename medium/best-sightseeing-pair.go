package main

import "fmt"

//固定fix
func maxScoreSightseeingPair(A []int) int {
	max := 0
	fix := A[0]
	for i := 1; i < len(A); i++ {
		max = maxFunc(max, fix+A[i]-i)
		fix = maxFunc(fix, A[i]+i)
	}
	return max
}

//暴力枚举
//func maxScoreSightseeingPair1(A []int) int {
//	max := 0
//	end := len(A) - 1
//	for i := 0; i < end; i++ {
//		for j := i + 1; j <= end; j++ {
//			score := A[i] + A[j] + i - j
//			if score >= max {
//				max = score
//			}
//		}
//	}
//	return max
//}

func maxFunc(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	A := []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(A))
}
