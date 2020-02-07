package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "ad be cf"
	fmt.Println(reverseWords(s))
}

// 编码最简单, 效率一般般
func reverseWords2(s string) string {
	res := strings.Split(s, " ")
	tmp := ""
	for i := 0; i < len(res)/2; i++ {
		tmp = res[len(res)-1-i]
		res[len(res)-1-i] = res[i]
		res[i] = tmp
	}
	tmp = res[0]
	for i := 1; i < len(res); i++ {
		if res[i] != "" {
			tmp = tmp + " " + res[i]
		}
	}
	return tmp
}

// 使用byte类型, 相同逻辑, 效率大大提升
func reverseWords(s string) string {
	// string换成[]byte 会大大提高效率
	srcBytes := []byte(s)
	// 分组,二位byte数组, 每一行是一个单词实际上
	bytesSlice := bytes.Split(srcBytes, []byte{' '})
	var dstBytes []byte
	for i := len(bytesSlice) - 1; i >= 0; i-- {
		// 剔除[]byte{''}的元素
		if len(bytesSlice[i]) == 0 {
			continue
		}
		dstBytes = append(dstBytes, bytesSlice[i]...)
		if i > 0 {
			dstBytes = append(dstBytes, ' ')
		}
	}

	l := len(dstBytes)
	// 判断dstBytes==[]byte{''}的情况
	if l != 0 {
		if dstBytes[l-1] == ' ' {
			dstBytes = dstBytes[:l-1]
		}
	}
	return string(dstBytes)
}
