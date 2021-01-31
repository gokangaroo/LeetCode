package main

import "fmt"

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}

func numSimilarGroups(strs []string) int {
	/*
	   问题: 多少个相似字符串组(a与b相似, b与c相似, abc为一组))
	   条件:
	       1. 都是小写单词
	       2. 字母组成都相同
	   如何判断相似:
	       除了2个字母顺序颠倒, 其他字母顺序一致
	*/
	// 双重循环+并查集, 主要是如何判断相似=>只有2个位置刚好不相等或者相等
	var uf = make([]int, len(strs))
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		for uf[i] != i {
			return find(uf[i])
		}
		return i
	}
	var union func(x, y int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX < rootY {
			uf[rootY] = rootX
		} else {
			uf[rootX] = rootY
		}
	}

	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if checkSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}

	var rootMap = make(map[int]bool)
	for i := 0; i < len(uf); i++ {
		rootMap[find(i)] = true
	}
	return len(rootMap)
}

func checkSimilar(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return count == 2 || count == 0
}
