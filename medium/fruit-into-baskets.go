package main

import (
	"fmt"
)

func main() {
	fruits := []int{5,9,0,9,6,9,6,9,9,9}
	fmt.Println(totalFruit(fruits))
}

func totalFruit(fruits []int) int {
	numCount := make(map[int]int, 0)
	var max,i,j int
	for i < len(fruits)-max {
//		fmt.Println(numCount)
		if len(numCount) <= 2 {
			numCount[fruits[j]] += 1
			j++
		}
		if len(numCount) <= 2 {
			if j-i > max {
				max = j - i
			}
		} else {
			numCount[fruits[i]] -= 1
			if numCount[fruits[i]] == 0 {
				delete(numCount, fruits[i])
			}
			i++
		}
	}
	return max
}
