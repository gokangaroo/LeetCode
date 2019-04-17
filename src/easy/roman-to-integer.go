package main

import "fmt"

func romanToInt(s string) int {
	data := []byte(s)
	res  := 0
	for k,v := range data{
		if k==0{
			res += character[v]
			continue
		}
		if data[k-1] == 'I'{
			if v=='V'||v=='X'{
				res = res-2*character[data[k-1]]
			}
		}
		if data[k-1] == 'X'{
			if v=='L'||v=='C'{
				res = res-2*character[data[k-1]]
			}
		}
		if data[k-1] == 'C'{
			if v=='D'||v=='M'{
				res = res-2*character[data[k-1]]
			}
		}
		res += character[v]
	}
	return res
}

var character = map[byte]int{
	'I':1,
	'V':5,
	'X':10,
	'L':50,
	'C':100,
	'D':500,
	'M':1000,
}

func main(){
	 str := "LVIII"
	 res := romanToInt(str)
	 fmt.Println(res)
}
