## LeetCode Golang 算法题

每周更新`LeetCode`题库集`Answers`

### 文件命名规则

题目名字中文`slug` 如 12. 整数转罗马数字即为 `integer-to-roman.go`

欢迎提交

### 解题思路
### algorithm programming
- 1.E: two-sum:两数之和

    在两个数组找到和为x的两个数的下标

    O(n), range数组, 判断targetNum(sum-rangeNum)是否已经在取出的map中
- 2.M: add-two-integers:计算两个逆序链表数的和
  
   计算两个逆序链表数的和
   
   就是要实现一下进位, 主要是用链表实现大数相加, 最终返回的链表也是逆序的
-  3.M: longest-substring-without-repeating-characters: 最长不重复子串
   
    最长不重复子串(获取长度)
    
    使用一个可截取的字符串,或者队列(FIFO)来遍历一遍原字符串,当有字母重复的时候,开始从头删,一直到删除重复的元素,然后继续遍历,最后O(n)得到max
- 4.H: median-of-two-sorted-arrays: 两个排序数组的中位数

    两个排序数组的中位数(合并后的中位数)
    
    这道题,主要是时间复杂度有要求O(log(m+n)), 先得到len,得到有利条奇偶数和谁长谁短(完全可以枚举所有场景,各个击破)
    
    奇数情况下,left多放一个,取left max,偶数情况,取(left max+right min)/2, 可以利用二分,快速得到临界值.比如取短的数组为m, 长的数组为n,m取在中间设为i,j就比较简单,i+j为m+n的一半即可
- 5.M: longest-palindromic-substring:  最长回文字符串(获取子串)

    最长回文字符串(获取子串)
    
    遍历原字符串, 写一个方法, 向字符两侧扩散, 第一次不相等就退出该方法, 但每一次都需要考虑两种情况, 一种是s[i]为核心扩散, 一种是s[i]=s[i+1]需要额外考虑两者为中心的情况.
- 16.M: 3sum-closest: 最接近的三数之和(一个int数组)

    指定一个数, 在一个数组内, 找到最接近这个数的三个数之和
    
    首先, 需要排序, 其次, 犁地式便利, i, 然后首尾left=i+1, right=len(n)-1, 计算sum与target的绝对值差距(根据summ>target判断动左还是右), 最终选出这个差距最小的为sum

- 17.M: letter-combinations-of-a-phone-number: 电话号码的字母组合(9键输入法)

    给几个数, 找到所有可能的字母组合.
    
    首先, 就是遍历, 排除1, 然后第一个数字的所有字母放到一个[]string里面, 然后与后面字母的[]string组合放到新的数组, 以此类推.
    
### databases sql
### shell
### multi threads