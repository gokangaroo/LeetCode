package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}

//双指针 头尾夹
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for {
		x := right - left
		y := minNum(height[left], height[right])
		if max < x*y {
			max = x * y
		}
		if left >= right {
			break
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return max
}

func minNum(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
