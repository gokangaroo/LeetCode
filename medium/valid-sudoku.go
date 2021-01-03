package main

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	// 1.9x9, [9][9]byte
	// 2.空格是.
	// 3.每一行, 每一列, 每一个3x3只出现一次
	// 简单思路: 通过x,y进行分组, 使用map判断是否重复出现
	// 九宫格组: "x/3y/3"字符串分组为key
	// 行列组: x和y分组为key
	// 后续优化: 1.因为实际每一组最多9个数, 使用数组也可以
	// 2. 优化到直接的分组规律, 而不是取巧的使用字符串填充
	// 3. for循环需要改为y先, x后, 充分利用cpu缓存
	// 4. 待定...
	groups := make(map[string]map[byte]bool)
	for x := 0; x < 9; x++ {
		gx := fmt.Sprintf("x%d", x)
		groups[gx] = make(map[byte]bool) // 列组
		for y := 0; y < 9; y++ {
			gy := fmt.Sprintf("y%d", y)
			if x == 0 {
				groups[gy] = make(map[byte]bool) // 行组
			}
			g := fmt.Sprintf("x%dy%d", x/3, y/3)
			if x%3 == 0 && y%3 == 0 {
				groups[g] = make(map[byte]bool) // 九宫格组
			}
			// 如果是. 跳过
			if board[x][y] == '.' {
				continue
			}
			// 判断是否存在
			if groups[g][board[x][y]] ||
				groups[gx][board[x][y]] ||
				groups[gy][board[x][y]] {
				return false
			}
			groups[g][board[x][y]] = true
			groups[gx][board[x][y]] = true
			groups[gy][board[x][y]] = true
		}
	}
	return true
}
