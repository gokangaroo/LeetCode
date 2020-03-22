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

func maxArea2(height []int) int {
	// 1.min(左,右)*(right-left)
	left, right := 0, len(height)-1
	max := min(height[left], height[right]) * (right - left)
	for left < right {
		// 2.贪心算法,缩短一格, 那高度去掉短的那个
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
		tmp := min(height[left], height[right]) * (right - left)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func min(left, right int) int {
	if left > right {
		return right
	}
	return left
}
