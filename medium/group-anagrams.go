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
	for _,v := range hash{
		res = append(res,v)
	}
	return res
}
