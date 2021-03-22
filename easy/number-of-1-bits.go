package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(0b11111111_11111111_11111111_11111101))
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
