package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

// 暴力四重循环
func restoreIpAddresses(s string) []string {
	// ip由4段组成, 我只要保证a+b+c+d=len(s)就行
	// 而且顺序是一定的, 更直接了.
	res := make([]string, 0)
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == len(s) {
						ia, _ := strconv.Atoi(s[:a])
						ib, _ := strconv.Atoi(s[a : a+b])
						ic, _ := strconv.Atoi(s[a+b : a+b+c])
						id, _ := strconv.Atoi(s[a+b+c:])
						if ia <= 255 && ib <= 255 && ic <= 255 && id <= 255 {
							tmp := fmt.Sprintf("%d.%d.%d.%d", ia, ib, ic, id)
							// 需要保证,0开头的只能是0本身
							if len(tmp) == len(s)+3 {
								res = append(res, tmp)
							}
						}
					}
				}
			}
		}
	}
	return res
}
