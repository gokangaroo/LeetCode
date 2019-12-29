package main

import "fmt"

var digitMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {
	digits := "29"
	arr := letterCombinations(digits)
	fmt.Println(arr)
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		if digits[i] == '1' {
			continue
		}
		//为空
		if len(res) == 0 {
			res = digitMap[digits[i]]
			continue
		}
		//组合
		tmpRes := make([]string, 0)
		for m := 0; m < len(res); m++ {
			for n := 0; n < len(digitMap[digits[i]]); n++ {
				tmpRes = append(tmpRes, res[m]+digitMap[digits[i]][n])
			}
		}
		res = tmpRes
	}
	return res
}
