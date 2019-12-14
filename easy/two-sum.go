package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)//创建map, key存值, value存索引(相同值会覆盖索引,也就是会默认选择最近的一个数)
	for i, value := range nums { //value=nums[i]
		mid, ok := mymap[target-value]//查看是否已经有放入的对应的value
		if ok {
			return []int{mid, i}//如果有,返回前面value索引和当前索引i
		}
		mymap[value] = i//如果没有将当前的遍历放入map
	}
	panic("没有两个数相加等于目标数")//遍历结束都没有, 那就是没有喽!
}
