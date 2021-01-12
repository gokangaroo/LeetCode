package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestStringWithSwaps("pwqlmqm",
		[][]int{{5, 3}, {3, 0}, {5, 1}, {1, 1}, {1, 5}, {3, 0}, {0, 2}}))
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// s按pairs交换, pairs每个交换都可以进行多次, 换成'最小的'
	// 等价于, 如果[a,b],a的字符>=b的字符, 就交换=>进一步理解成排序
	// [a,b]和[b,c]可以理解成[a,b,c]的排序=>进一步可以理解成3位, 4位的排序
	n := len(s)
	ds := newDjset(n)
	for _, v := range pairs {
		ds.Merge(v[0], v[1])
	}
	nps := make(map[int][]int)
	for i := 0; i < n; i++ {
		nps[ds.Find(i)] = append(nps[ds.Find(i)], i)
	}

	bs := []byte(s)
	for _, pair := range nps {
		smallestStringWithSort(bs, pair)
	}
	return string(bs)
}

func smallestStringWithSort(s []byte, pair []int) {
	sort.Ints(pair)

	var tmp = make([]byte, len(pair))
	for i, _ := range pair {
		tmp[i] = s[pair[i]]
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	for i, _ := range pair {
		s[pair[i]] = tmp[i]
	}
}

type Djset struct {
	Parent []int
	Rank   []int
}

func newDjset(n int) Djset {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return Djset{parent, rank}
}

func (ds Djset) Find(x int) int {
	if ds.Parent[x] != x {
		ds.Parent[x] = ds.Find(ds.Parent[x])
	}
	return ds.Parent[x]
}

func (ds Djset) Merge(x, y int) {
	rx := ds.Find(x)
	ry := ds.Find(y)
	if rx == ry {
		return
	}

	if ds.Rank[rx] < ds.Rank[ry] {
		rx, ry = ry, rx
	}
	ds.Parent[ry] = rx // ry < rx
	if ds.Rank[rx] == ds.Rank[ry] {
		ds.Rank[rx] += 1
	}
}
