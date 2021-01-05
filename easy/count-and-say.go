package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

const byteStart = 48

func countAndSay(n int) string {
	// 描述countAndSay(n-1)的string
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	x := countAndSay(n - 1)
	// 优化内存分配
	// TODO 还能进一步优化缓存结果集
	res := make([]byte, 0, 2*len(x))
	var count byte = 1
	for i := 1; i < len(x); i++ {
		if x[i] == x[i-1] {
			count++
		} else {
			// do append & new start
			res = append(res, []byte{count + byteStart, x[i-1]}...)
			count = 1
		}
	}
	res = append(res, []byte{count + byteStart, x[len(x)-1]}...)
	return string(res)
}
