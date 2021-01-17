package main

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 1}, {1, 2}, {2, 1}}))
}

func checkStraightLine(coordinates [][]int) bool {
	// 判断在一条线,点都不重复
	// 2<=len<=1000
	// x,y 都是+-1w, 所以乘法不会溢出, 乘法可以不用考虑分母为0的场景
	xWidth := coordinates[1][0] - coordinates[0][0]
	// if xWidth==0{//说明是垂直的线
	//     x:=coordinates[0][0]
	//     for i:=1;i<len(coordinates);i++{
	//         if coordinates[i][0]!=x{
	//             return false
	//         }
	//     }
	//     return true
	// }
	yHeight := coordinates[1][1] - coordinates[0][1]
	// if yHeight==0{//说明是水平的线
	//     y:=coordinates[0][1]
	//     for i:=1;i<len(coordinates);i++{
	//         if coordinates[i][1]!=y{
	//             return false
	//         }
	//     }
	//     return true
	// }
	// 斜着的线 yHeight2*xWidth=xWidth2*yHeight
	for i := 1; i < len(coordinates); i++ {
		xWidth2 := coordinates[i][0] - coordinates[0][0]
		yHeight2 := coordinates[i][1] - coordinates[0][1]
		if yHeight2*xWidth != xWidth2*yHeight {
			return false
		}
	}
	return true
}
