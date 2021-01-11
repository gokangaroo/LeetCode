package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// s按pairs交换, pairs每个交换都可以进行多次, 换成'最小的'
	// 等价于, 如果[a,b],a的字符>=b的字符, 就交换=>进一步理解成排序
	// [a,b]和[b,c]可以理解成[a,b,c]的排序=>进一步可以理解成3位, 4位的排序
	bs := strings.Split(s, "")
	pairs = unionPairs(pairs, len(bs))
	for _, pair := range pairs {
		smallestStringWithSort(bs, pair)
	}
	return strings.Join(bs, "")
}

func smallestStringWithSort(s []string, pair []int) {
	sort.Ints(pair)

	var tmp = make([]string, len(pair))
	for i, _ := range pair {
		tmp[i] = s[pair[i]]
	}
	sort.Strings(tmp)

	for i, _ := range pair {
		s[pair[i]] = tmp[i]
	}
}

func unionPairs(pairs [][]int, l int) (res [][]int) {
	// 合并有交集的所有pairs
	// TODO 并查集, 父子节点
	var root = make([]int, l) // idx:int->root
	for i, _ := range pairs {
		sort.Ints(pairs[i])
		var r1 = pairs[i][0]
		var r2 = pairs[i][1]
		//for {
		// TODO solve zero
		root[r1] = r1
		root[r2] = r1
		//}
	}
	var rootSlice = make(map[int][]int, 0)
	for i, v := range root { // int->root
		if rootSlice[v] == nil {
			rootSlice[v] = []int{v}
		}
		rootSlice[v] = append(rootSlice[v], i)
	}
	for i, _ := range rootSlice {
		res = append(res, rootSlice[i])
	}
	return res
}
