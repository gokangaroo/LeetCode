package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
}

// 判断单调函数, > < =,其中=表示存在于两者, 使用计数逻辑简单清晰
func isMonotonic(A []int) bool {
	var asc, desc int
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			if desc > asc {
				return false
			}
			asc++
		} else if A[i] > A[i+1] {
			if asc > desc {
				return false
			}
			desc++
		} else {
			asc++
			desc++
		}
	}
	return asc == len(A)-1 || desc == len(A)-1
}
