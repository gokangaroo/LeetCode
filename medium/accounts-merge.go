package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}))
}

func accountsMerge(accounts [][]string) [][]string {
	/*
	   1.有相同邮箱的进行合并, v[0]是name, v[1:]...是邮箱
	   2.合并后返回, v[1:]需要排序: sort.Strings(v[1:])
	   3.相同名字的不一定是同一个人, 但是不同名字的必定不是一个人
	   使用并查集, 使用accounts下标
	*/
	// 并查集1: parents设为自己
	parents := make([]int, len(accounts))
	for i, _ := range parents {
		parents[i] = i
	}

	// 并查集2: find parents
	var find func(idx int) int
	find = func(idx int) int {
		for parents[idx] != idx {
			return find(parents[idx])
		}
		return idx
	}

	// 并查集3: union x's root & y's root
	var union func(x, y int)
	union = func(x, y int) {
		rootx := find(x)
		rooty := find(y)
		parents[rootx] = parents[rooty]
	}

	// 并查集4: start union
	var mmap = make(map[string]int) // email:idx
	for i, v := range accounts {
		for j := 1; j < len(v); j++ {
			if idx, ok := mmap[v[j]]; ok {
				union(i, idx)
			} else {
				mmap[v[j]] = i
			}
		}
	}

	// 并查集5: 产出结果
	res := make([][]string, 0, len(accounts))
	mmap2 := make(map[int]int) // idx=>new idx
	for email, idx := range mmap {
		p := find(idx)
		if nIdx, ok := mmap2[p]; !ok {
			res = append(res, []string{accounts[p][0], email})
			mmap2[p] = len(res) - 1 //记得赋值
		} else {
			res[nIdx] = append(res[nIdx], email)
		}
	}

	for i := 0; i < len(res); i++ {
		sort.Strings(res[i][1:])
	}

	return res
}
