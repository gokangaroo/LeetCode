package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 = "123456789364732432648632482374217472", "9876543214t2384238747234792749274927374"
	res := multiply(num1, num2)
	fmt.Println(res)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	//每轮结果字符串
	baseStr := []string{"0"}

	lenIndex := len(num2)
	for i := 1; i <= lenIndex; i++ {
		//附加数
		tmpExtra := 0
		str := make([]string, 0)
		for j := len(num1) - 1; j >= 0; j-- {
			tmpNum1, _ := strconv.Atoi(string(num1[j]))
			tmpNum2, _ := strconv.Atoi(string(num2[lenIndex-i]))
			tmpNum := tmpNum1 * tmpNum2
			m := (tmpNum + tmpExtra) / 10
			n := (tmpNum + tmpExtra) % 10
			str = append(str, strconv.Itoa(n))
			//第一个数字大于10
			if j == 0 && m > 0 {
				str = append(str, strconv.Itoa(m))
			}
			if m > 0 {
				tmpExtra = m
			} else {
				tmpExtra = 0
			}
		}
		baseStr = addTwoString(baseStr, str, i-1)
	}

	res := ""
	for i := len(baseStr) - 1; i >= 0; i-- {
		res += baseStr[i]
	}
	return res
}

//两个字符串根据一定间隔位数相加
func addTwoString(baseStr []string, str []string, index int) []string {
	res := make([]string, 0)
	tmpExtra := 0
	for i := 0; i < len(baseStr); i++ {
		num1, _ := strconv.Atoi(baseStr[i])
		num2 := 0
		//不超过str长度
		if i-index >= 0 {
			num2, _ = strconv.Atoi(str[i-index])
		}
		sum := num1 + num2
		if sum+tmpExtra >= 10 {
			m := (sum + tmpExtra) % 10
			res = append(res, strconv.Itoa(m))
			tmpExtra = 1
		} else {
			res = append(res, strconv.Itoa(sum+tmpExtra))
			tmpExtra = 0
		}

		//超过baseStr长度
		if i == len(baseStr)-1 {
			for {
				if i-index+1 == len(str) {
					if tmpExtra > 0 {
						res = append(res, "1")
					}
					break
				}
				num2, _ = strconv.Atoi(str[i-index+1])
				if num2+tmpExtra == 10 {
					res = append(res, "0")
					tmpExtra = 1
				} else {
					res = append(res, strconv.Itoa(num2+tmpExtra))
					tmpExtra = 0
				}
				i++
			}
		}

	}
	return res
}