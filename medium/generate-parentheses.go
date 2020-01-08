package main

import "fmt"

func main() {
	n := 2
	res := generateParenthesis(n)
	fmt.Println(res)

}

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	dfs("", 0, 0, n, &res)
	return res
}

func dfs(str string, left int, right int, n int, res *[]string){
	//括号都使用完了  结果结算
	if left == n && right == n {
		*res = append(*res, str)
		return
	}

	//左括号不会比右括号少 回溯
	if left < right {
		return
	}

	//进入左子树
	if left < n {
		dfs(str+"(", left+1, right, n, res)
	}
	//进入右子树
	if right < n {
		dfs(str+")", left, right+1, n, res)
	}
}
