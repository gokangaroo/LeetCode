package main

import "fmt"

func main() {
	is := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	fmt.Println(len(is))
	by5 := prefixesDivBy5(is)
	fmt.Println(len(by5))
	fmt.Println(by5)
}

func prefixesDivBy5(A []int) []bool {
	var (
		res = make([]bool, 0, len(A))
		num = 0
	)
	for i := 0; i < len(A); i++ {
		num = (num*2 + A[i]) % 5 // 余掉的数, 下次还是会余掉, 所以只要留余数就好了
		res = append(res, num == 0)
	}
	return res
}
