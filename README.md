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
- 2.M: add-two-numbers:计算两个逆序链表数的和
  
   计算两个逆序链表数的和
   
   就是要实现一下进位, 主要是用链表实现大数相加, 最终返回的链表(head+cursor)也是逆序的, 最后要检查下有没有遗留进位.
   
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

- 11.M: Container With Most Water : 盛最多水的容器

    ![](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)
    
    就是一个数组, 选出两个值, 以x轴为底部, 如何乘最多的水(短的那一根*底部宽度)<br/>
    双指针从两头开始, 然后面临left和right哪个移动的问题, 这里使用贪心算法, 短的那一根舍弃掉.

- 15.M: 3sum: 三数之和

    该题有十分多的重复数字, 而之前twosum只要找到一个就好, 在这里行不通, 需要先排序.<br/>
    然后假定有left,mid,right三个指针, 这时候最好的选择是以1撬动2, 最好是left撬动mid&right, 当left>0,结束<br/>
    left取非重复的第一个, mid取left+1, right固定从结尾开始, 对于一个left可能有多个mid和right, 所以for mid<right, <br/>
    mid和right取到非重复的数, 继续循环, 当mid>right表示这个left已经没有结果了, 回到前面.
    > 可以把mid写成left, left写成i, 比如下面的16

- 16.M: 3sum-closest: 最接近的三数之和(一个int数组)

    指定一个数, 在一个数组内, 找到最接近这个数的三个数之和
    
    首先, 需要排序, 其次, 犁地式便利, i, 然后首尾left=i+1, right=len(n)-1, 计算sum与target的绝对值差距(根据sum>target判断动左还是右), 最终选出这个差距最小的sum<br>
    注意点: 如果找到相等的, 可以提前退出; 这题不用在意重复值(无需增加编码复杂度).

- 17.M: letter-combinations-of-a-phone-number: 电话号码的字母组合(9键输入法)

    给几个数, 找到所有可能的字母组合.
    
    首先, 就是遍历, 排除1, 然后第一个数字的所有字母放到一个[]string里面, 然后与后面字母的[]string组合放到新的数组, 以此类推.
    
- 22.M: generate-parentheses: 括号生成

    给一个n, 给出n对括号组合的所有情况,括号顺序一定要对.
    
    这一道题有很多方法, 深度遍历, 广度遍历, 还可以使用动态规划.
    
- 71.M: simplify-path: 简化路径

    以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径
    
    这个是栈的经典题 你可以有多个选择 以单个字符做校验还有就是按照目录 这里最合适以/的目录划分 上级目录对应着出栈 其余视为入栈或者跳过即可
   
- 91.M: Decode Ways: 解码方法(这个题目0好碍事,类似走楼梯, 可以使用动态规划.)
    
    如果s[i]是0, 那么s[i-1]必须是1和2,不然就return 0,而且dp[i-1]无效,dp[i]=dp[i-2]==>i-1和i必须一起<br/>
    然后ParseInt(s[i-1:i+1],10,64),(当然也可以判断i和i-1两个字符)如果10< <=26,说明可以合并分开两种编码, dp[i]=dp[i-1]+dp[i-2]<br/>
    如果<10, 说明只能自己编码, dp[i]=dp[i-1]
   
- 102.M: Binary Tree Level Order Traversal: 二叉树层次遍历

    利用前序遍历, 每一个递归调用都让depth+1(插入到指定深度的数组)
    > res必须是公用, 每次开局初始化一下, 如547那边的M实际是系统提供的公用二维数组.
   
- 173.M: Binary Search Tree Iterator: 二叉搜索树迭代器
    
    调用Next可以不断弹出二叉搜索树内最小的值,NextHash返回bool
   
    1. 中序遍历二叉树(递归), 初始化构建为一个数组, 假装是二叉搜索树 (根据index判断hashNext, 默认-1)
    
    2. 自建一个栈, 中序遍历, 添加自己`并加入左节点`(左节点在上), (右边节点在弹出时,使用后加入).
   
- 215.M: kth-largest-element-in-an-array: 数组中的第K个最大元素
    
    在未排序的数组中找到第 k 个最大的元素
    
    大顶堆:采用大堆排序(节点不小于其子节点) 第几个最大值也就是第几次建堆 这样可以得出第几个最大值 <br/>
    小顶堆:先k-1个数形成小顶堆,额外记录一个min, 就是kth,遍历剩余元素, 如果>kth, 先push小顶堆,kth=pop, (pop必定大于原kth)

- 300.M: longest-increasing-subsequence: 最长上升子序列
    
    动态规划: d[i]表示以nums[i]结尾的最长子序列的长度, j从0-(i-1), <br/>
    找出最大的,满足nums[j]<nums[i]的d[j],d[i]=d[j]+1(可以接在后面), 如果d[j]>max, 则替换max
    
    贪心算法+二分法替~~~~换(利用left): tail数组, 表示已经找到的最大升序列, 在遍历nums的时候, <br/>
    如果nums[i]比tail的最后一个已插入值大, 那么就直接接到后面.如果比最后一个小, 可以直接替换吗?<br/>
    不能!只能根据二分法替换tail中第一个大于自己的数, 如果恰好是最后一个, 则完成了贪心, 让最后的高度变矮了.<br/>
    而如果换在中间, 则是一个无关紧要的贪心, 后续有值找插入的时候, 对其不会造成影响, 对结果也没什么影响.

- 547.M:  Friend Circles:朋友圈
    
    M[i][j]=1表示i与j是朋友, M[i][i]初始都是1, 找朋友圈的个数
    
    就是一个深度遍历, 如果M[i][i]还是1, dfs(M,i),count++ <br/>
    然后染色,M[i][j]==1&&M[j][j]!=-1, dfs(M,j);M[j][j]=-1 <br/>
    最后的话, count就是朋友圈个数

- 567.M:  checkInclusion: 字符串的排列
    两个字符串s1, 和s2, 判断s2是否有s1的排列之一作为子串(这个排列可太多了): 只有小写
    
    首先,要缩小范围, len(s1), 对s1进行计数, 然后双指针遍历s2, left和right间隔len<br/>
    left开始前进(right<len(s2)+1), 如果计数不够用, 如果right-left==l, 就left和right一起平移<br/>
    否则就right-len开始进行剔除(判断cursor是否==left), 直到count足够, right再往后退<br/>
    如果right到头, left还失败了, 那就返回false.
    
- 622.M: Design Circular Queue:设计循环队列
    
    就是循环空间的队列, 满了就不能插入, 可以类比redolog
    
    使用数组, head和tail记录头尾, 当head=tail的时候, 队列已满,或者已经清空<br/>
    添加两个字段直接表示这个.当出现事件时修改.(目前代码没有加Mutex,实际如果需要一定要加)


### databases sql
### shell
### multi threads