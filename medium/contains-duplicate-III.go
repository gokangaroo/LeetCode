package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-3, 3}
	k := 2
	t := 4
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

// 先使用n^2解法看看: 时间复杂度O(n*min(k,n)), 空间复杂度O(1)
// 存在i,j满足条件
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	max := func(k, n int) int {
		if k > n {
			return k
		}
		return n
	}
	for i := 0; i < len(nums); i++ { //需要考虑k>len的场景
		for j := max(i-k, 0); j < i; j++ {
			if !(nums[j]-nums[i] < - 1*t || nums[j]-nums[i] > t) {
				return true
			}
		}
	}
	return false
}

// 利用排序特性,二分查找=>TODO 也可以改成二叉搜索树版本, 进行删除以及插入(下面切片方式已经优化到了极致)
// 时间复杂度O(n*log(k)), 空间复杂度O(k)
// 进一步: 如果k>n,实际分配n即可: 时间复杂度O(n*log(min(n,k))), 空间复杂度O(min(n,k))
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	min := func(k, n int) int {
		if k > n {
			return n
		}
		return k
	}
	tmp := make(sortInts, 0, min(len(nums), k+1))
	for i := 0; i < len(nums); i++ { //不能是len(nums)-k, 因为可能k更大
		if tmp.Len() == k+1 {
			//应该删掉开头那个数
			tmp.delXInSlice(nums[i-k-1])
		}
		if tmp.addXAndCompare(nums[i], t) {
			return true
		}
	}
	return false
}

// 利用排序特性, 比较最大最小值即可
type sortInts []int

func (s *sortInts) Len() int {
	return len(*s)
}

// 删除某个值
func (s *sortInts) delXInSlice(x int) {
	left := 0
	right := s.Len() - 1
	for left <= right {
		mid := (left + right) / 2
		if (*s)[mid] == x { //进行剔除
			if mid == len(*s)-1 {
				*s = (*s)[:mid]
			} else {
				*s = append((*s)[:mid], (*s)[mid+1:]...)
			}
			return
		}
		if (*s)[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

// 添加某个值并判断是否满足条件
func (s *sortInts) addXAndCompare(x, t int) bool {
	left := 0
	right := s.Len() - 1
	if s.Len() == 0 {
		*s = append(*s, x)
		return false
	}
	if x >= (*s)[right] {
		*s = append(*s, x)
		return x-(*s)[right] <= t
	}
	if x <= (*s)[left] {
		*s = append([]int{x}, *s...)
		return (*s)[1]-x <= t
	}
	for left <= right {
		mid := (left + right) / 2 // mid表示小于x的最大的数
		if (*s)[mid] <= x && (*s)[mid+1] >= x {
			if x-(*s)[mid] <= t || (*s)[mid+1]-x <= t {
				return true
			} else {
				// 这一步需要copy
				//dst := make([]int, mid+1)
				//copy(dst, (*s)[:mid+1])
				//*s = append(append(dst, x), (*s)[mid+1:]...)
				// 理解切片后, 可以减少内存使用并大幅度减少拷贝耗时.
				*s = append(*s, (*s)[s.Len()-1])
				copy((*s)[mid+2:], (*s)[mid+1:]) //后移
				(*s)[mid+1] = x
				return false
			}
		}
		if (*s)[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// 最快的桶方法/map方法
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	res := make(mineMap)
	if t < 0 {
		return false
	}
	for i := 0; i < len(nums); i++ { //i 是后数的下标
		var idx int
		if t == 0 {
			idx = nums[i]
		} else {
			idx = nums[i] / t //idx是k
		}
		for bucketIdx := idx - 1; bucketIdx < idx+2; bucketIdx++ {
			if x, ok := res[bucketIdx]; ok {
				for j := len(x) - 1; j >= 0; j-- { //x[j]是前数的下标
					if i-x[j] > k {
						break
					} else {
						if int(math.Abs(float64(nums[i]-nums[x[j]]))) <= t {
							return true
						}
					}
				}
			}
		}
		//res[idx] = append(res[idx], i) //存的是下标
		//sort.Ints(res[idx])
		//利用二分法提高性能,感觉没有必要, t不大的情况下提升不大
		//addNewNum(i, &res[idx])//map内的v是无法取得地址的, 因为map会进行rehash, 会产生悬挂指针: https://github.com/golang/go/issues/11865
		aim := res[idx]
		addNewNum(i, &aim)
		res[idx] = aim
	}
	return false
}

func addNewNum(x int, s *[]int) {
	left := 0
	right := len(*s) - 1
	if len(*s) == 0 {
		*s = append(*s, x)
		return
	}
	if x >= (*s)[right] {
		*s = append(*s, x)
		return
	}
	if x <= (*s)[left] {
		*s = append([]int{x}, *s...)
		return
	}
	for left <= right {
		mid := (left + right) / 2 // mid表示小于x的最大的数
		if (*s)[mid] <= x && (*s)[mid+1] >= x {
			*s = append(*s, (*s)[len(*s)-1])
			copy((*s)[mid+2:], (*s)[mid+1:]) //后移
			(*s)[mid+1] = x
			return
		}
		if (*s)[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

type mineMap map[int][]int //存入 的不是数值而是idx
