package main

import "fmt"

func main() {
	fmt.Println(clumsy(10))
	fmt.Println(clumsyFool(10))
}

func clumsy(N int) int {
	/*
	   首先想到的是栈, 如果是*和/就将栈顶取出计算, +入栈, -换成负数入栈
	*/
	stack := make([]int, 0, N/2+1)
	// chars:=[]byte{'*','/','+','-'}
	stack = append(stack, N)
	for i := N - 1; i > 0; i-- {
		switch (N - i - 1) % 4 { // 当前处于chars的下标
		case 0:
			stack[len(stack)-1] *= i
		case 1:
			stack[len(stack)-1] /= i
		case 2:
			stack = append(stack, i)
		case 3:
			stack = append(stack, -1*i)
		}
	}
	r := 0
	for i := 0; i < len(stack); i++ {
		r += stack[i]
	}
	return r
}

func clumsyFool(N int) (ans int) {
	switch {
	case N == 1:
		return 1
	case N == 2:
		return 2
	case N == 3:
		return 6
	case N == 4:
		return 7

	case N%4 == 0:
		return N + 1
	case N%4 <= 2:
		return N + 2
	default:
		return N - 1
	}
}
