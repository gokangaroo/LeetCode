package main

import "fmt"

func main() {
	isConnected := [][]int{{1}}
	//isConnected := [][]int{
	//	{1, 1, 0},
	//	{1, 1, 0},
	//	{0, 0, 1}}
	fmt.Println(findCircleNumCity(isConnected))
}

func findCircleNumCity(isConnected [][]int) int {
	n := len(isConnected)
	if n == 1 {
		return 1
	}

	us := generate(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isConnected[i][j] == 1 {
				us.union(i, j)
			}
		}
	}
	return us.count
}

type unionSet struct {
	count    int
	parentOf []int
}

func generate(n int) (us *unionSet) {
	us = &unionSet{}
	us.count = n
	us.parentOf = make([]int, n)
	for i := 0; i < n; i++ {
		us.parentOf[i] = i //1. 初始化节点(祖先)都是自身
	}
	return
}

func (us *unionSet) union(p, q int) {
	if p == q {
		return
	}

	rootP := us.find(p)
	rootQ := us.find(q)
	if rootP == rootQ {
		return
	}

	us.parentOf[rootP] = rootQ // 2. 把p的节点改到q
	us.count--
}

func (us *unionSet) find(t int) (root int) {
	root = t
	for root != us.parentOf[root] { // 3. 循环找到根节点(祖先)
		root = us.parentOf[root]
	}
	us.parentOf[t] = root // 4. 懒优化, 查过一次直接指向根节点
	return
}
