## LeetCode Golang 算法题

每周更新`LeetCode`题库集`Answers`

### 文件命名规则

题目名字中文`slug` 如 12. 整数转罗马数字即为 `integer-to-roman.go`

欢迎提交

### 大概思路
1. e:two-sum:两数之和(在两个数组找到和为x的两个数的下标): O(n), range数组, 判断targetNum(sum-rangeNum)是否已经在取出的map中
2. m:add-two-integers:两数相加(计算两个逆序链表数的和): 就是要实现一下进位, 主要是用链表实现大数相加, 最终返回的链表也是逆序的
3. m:longest-substring-without-repeating-characters: 最长不重复子串(获取长度): 使用一个可截取的字符串,或者队列(FIFO)来遍历一遍,O(n)得到max