package sort_and_search

import (
	"fmt"
	"sort"
)

func Main() {
	nums1 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	nums2 := []int{-1, 1, 1, 1, 2, 3}

	merge(nums1, 3, nums2, 6)
	fmt.Println(nums1)
}

/*
三种方法：
1. 合并后排序
2. 双指针（复制到临时数组）
3. 逆向双指针（空间复杂度最低）
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	res := make([]int, m+n)
	for i < m && j < n {
		fmt.Println(i, j)
		if nums1[i] == nums2[j] {
			res[i+j], res[i+j+1] = nums1[i], nums1[i]
			i++
			j++
			continue
		}
		if nums1[i] > nums2[j] {
			res[i+j] = nums2[j]
			j++
		} else {
			res[i+j] = nums1[i]
			i++
		}
	}
	fmt.Println(i, j)
	if i < m {
		for ; i < m; i++ {
			res[i+j] = nums1[i]
		}
	}
	if j < n {
		for ; j < n; j++ {
			res[i+j] = nums2[j]
		}
	}
	copy(nums1, res)
}

var isBadVersion func(int) bool

// 用二分查找，但是sort包里面有非常适合用在这里的函数
func firstBadVersion(n int) int {
	return sort.Search(n, isBadVersion)
}
