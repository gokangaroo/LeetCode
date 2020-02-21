package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Print(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	strMap := make(map[string]int)
	res := make([]string, 0)

	for i := 0; i <= len(s)-10; i++ {
		if v, _ := strMap[s[i:i+10]]; v == 1 {
			res = append(res, s[i:i+10])
		}
		strMap[s[i:i+10]]++
	}
	return res
}

func findRepeatedDnaSequences1(s string) []string {
	sByte := []byte(s)
	length := len(sByte)
	strMap := make(map[int][]byte)
	res := make([]string, 0)

	if length <= 10 {
		return res
	}

	for i := 0; i <= length-10; i++ {
		if i == 0 {
			strMap[i] = sByte[i : i+10]
			continue
		}
		count := 0
		index := 0
		for k, v := range strMap {
			if bytes.Equal(v, sByte[i:i+10]) {
				index = k
				count++
				if count > 2 {
					break
				}
			}
		}
		if count == 1 {
			res = append(res, string(strMap[index]))
		}
		strMap[i] = sByte[i : i+10]
	}
	return res
}
