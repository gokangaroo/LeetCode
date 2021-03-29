package main

import "fmt"

func main() {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}

func reverseBits(num uint32) uint32 {
	var aim uint32 = 0
	for i := 0; i < 32; i++ {
		aim <<= 1
		if num&1 == 1 {
			aim += 1
		}
		num >>= 1
	}
	return aim
}
