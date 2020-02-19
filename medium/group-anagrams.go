package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{
		"eat",
		"tea",
		"tan",
		"ate",
		"nat",
		"bat",
	}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	hash := make(map[string][]string)

	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i int, j int) bool {
			return str[i] < str[j]
		})
		hash[string(str)] = append(hash[string(str)], v)
	}
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}

func groupAnagrams_o(strs []string) [][]string {
	// 使用质数乘积的方法.
	letters := map[byte]int{
		'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41,
		'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
	}
	res := make([][]string, 0)
	m := make(map[int]int, 0)
	l := len(strs)
	if l == 0 {
		return res
	}
	for i := 0; i < l; i++ {
		mt := 1
		tmp := []byte(strs[i])
		for _, v := range tmp {
			mt *= letters[v]
		}
		if v, ok := m[mt]; ok {
			res[v] = append(res[v], strs[i])
		} else {
			m[mt] = len(m)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}
