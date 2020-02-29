package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{
		[]int{5, 40},
		[]int{6, 8},
		[]int{6, 70},
		[]int{7, 9},
		[]int{2, 30},
		[]int{9, 11},
		[]int{13, 12},
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})
	fmt.Println(maxEnvelopes(envelopes))
}

//动态规划思想  不过子问题的解决需要往前一次比较 双子序列
func maxEnvelopes(envelopes [][]int) int {
	max := 0
	nr := len(envelopes)
	if nr == 1 || nr == 0 {
		return nr
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	arr := make([]int, nr)
	for i := 0; i < nr-1; i++ {
		//除了这种情况 其余全部回头查
		if envelopes[i+1][1] > envelopes[i][1] && envelopes[i+1][0] > envelopes[i][0] && max == arr[i] {
			arr[i+1] = arr[i] + 1
		} else {
			index := i
			//回头找上一个最大值
			k := 0
			tmpMax := -1
			for {
				if index-k >= 0 {
					if (envelopes[index-k][1] < envelopes[i+1][1]) && (envelopes[index-k][0] < envelopes[i+1][0]) {
						if tmpMax < arr[index-k] {
							tmpMax = arr[index-k]
						}
					}
					arr[i+1] = tmpMax + 1
				} else {
					break
				}
				k++
			}
		}
		if max < arr[i+1] {
			max = arr[i+1]
		}
	}
	return max + 1
}
