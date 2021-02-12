package main

import (
	"fmt"
)

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion0(s1 string, s2 string) bool {
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

// 计数法
func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)             //[0:l]
	count := make([]int, 26) // 如果是多字符就255
	for _, v := range s1 {
		count[int(v-'a')]++
	}
	left, right := 0, l
	for left < right && right < len(s2)+1 {
		if count[s2[left]-'a'] > 0 {
			count[s2[left]-'a']--
			left++
			continue
		}
		if right-l == left { //1.双指针平移动
			left++
			right++
			continue
		}
		// 2.开始推进
		cursor := right - l
		for s2[cursor] != s2[left] {
			count[s2[cursor]-'a']++
			cursor++
			right++
		}
		if cursor == left { //双指针平移动
			left++
			right++
		} else { //算是成功一个
			left++
			right++ //这里并不是随着left++,而是舍弃了cursor,虽然代码一样, 但是含义完全不同
		}
	}
	if left == right-1 { //鹊桥相会,说明是有这个序列.
		return true
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	/*
	   1. 对s1进行计数
	   2. 滑动窗口判断是否存在
	*/
	if len(s2) < len(s1) {
		return false
	}

	length := len(s1)
	count := make([]int, 26)
	misc := 0 // 表示还未"核平"的字母, 当misc=0时表示成功
	for _, v := range s1 {
		if count[v-'a'] == 0 {
			misc++
		}
		count[v-'a']++
	}

	for i := 0; i < length; i++ {
		count[s2[i]-'a']--
		if count[s2[i]-'a'] == -1 {
			misc++
		}
		if count[s2[i]-'a'] == 0 {
			misc--
		}
	}
	if misc == 0 {
		return true
	}

	for i := length; i < len(s2); i++ {
		// 左删 右增 i-length i
		count[s2[i-length]-'a']++
		if count[s2[i-length]-'a'] == 1 {
			misc++
		}
		if count[s2[i-length]-'a'] == 0 {
			misc--
		}
		count[s2[i]-'a']--
		if count[s2[i]-'a'] == -1 {
			misc++
		}
		if count[s2[i]-'a'] == 0 {
			misc--
		}
		if misc == 0 {
			return true
		}
	}

	return false
}
