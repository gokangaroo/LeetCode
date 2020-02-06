package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := "a-"
	s2 := "b0-1-a"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	count := make([]int, 256) //256个字符编码: 0-255, ascii码
	for _, v := range s1 {
		count[v]++
	}
	l, r := 0, 0
	for count[s2[r]] == 0 && r < len(s2)-1 { //找到s2中第一个s1中存在的字符
		l++
		r++
	}
	for r < len(s2) {
		if count[s2[r]] > 0 {
			count[s2[r]]--
			r++
		} else { //这个字符s1不存在,或者之前被消耗完了
			if r-l == len(s1) {
				return true
			}
			//把l放到与r第一个同字母的位置继续
			for s2[l] != s2[r] && l < r {
				count[s2[l]]++
				l++
			}
			if l == r { //说明是不存在,直接看下一个字符
				l++
				r++
			} else { //说明是一个重复点
				count[s2[l]]++
				l++
			}
		}
	}
	if r-l == len(s1) {
		return true
	}
	return false
}

// 完美超时, 应对短的字符串还可以.
func checkInclusion0(s1 string, s2 string) bool {
	m := len(s1)
	n := len(s2)
	if m > n {
		return false
	}
	// 排序法, 找s1的所有集合显然不合理
	// 那就不断截取s2与s1相同长度出来, 然后转为排序,最终比较'
	xx := []byte(s1)
	sort.Slice(xx, func(i, j int) bool {
		return xx[i] > xx[j]
	})
	for i := 0; i+m < n+1; i++ { //i+m可以到n, 而不是n-1
		hh := []byte(s2[i : i+m])
		sort.Slice(hh, func(i, j int) bool {
			return hh[i] > hh[j]
		})
		if string(hh) == string(xx) {
			return true
		}
	}
	return false
}
