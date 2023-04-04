package arrays

import (
	"fmt"
	"sort"
)

func Main() {
	board := [][]string{
		{".", ".", ".", ".", "5", ".", ".", "1", "."},
		{".", "4", ".", "3", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "3", ".", ".", "1"},
		{"8", ".", ".", ".", ".", ".", ".", "2", "."},
		{".", ".", "2", ".", "7", ".", ".", ".", "."},
		{".", "1", "5", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "2", ".", "9", ".", ".", ".", ".", "."},
		{".", ".", "4", ".", ".", ".", ".", ".", "."},
	}
	board2 := make([][]byte, 9)
	for i := range board2 {
		board2[i] = make([]byte, 9)
		for j, v := range board[i] {
			board2[i][j] = byte(v[0])
		}
	}

	fmt.Println(isValidSudoku(board2))
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	j := 0
	for i := 1; i < n; i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}

func maxProfit(prices []int) int {
	n := len(prices)
	j := 0
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			continue
		}
		if prices[j] >= prices[i-1] {
			j = i
			continue
		}
		profit += prices[i-1] - prices[j]
		j = i
	}
	if j != n-1 {
		profit += prices[n-1] - prices[j]
	}

	return profit
}

func rotate(nums []int, k int) {
	n := len(nums)
	if k > n {
		k %= n
	}
	copied := make([]int, n)
	for i := range nums {
		if i < k {
			copied[i] = nums[n-k+i]
		} else {
			copied[i] = nums[i-k]
		}
	}
	copy(nums, copied)
}

func containsDuplicate(nums []int) bool {
	n := len(nums)
	sort.Ints(nums)

	j := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[j] {
			return true
		}
		j = i
	}
	return false
}

func containsDuplicateV2(nums []int) bool {
	mp := make(map[int]bool)
	for _, v := range nums {
		_, prs := mp[v]
		if !prs {
			mp[v] = true
		} else {
			return true
		}
	}
	return false
}

func singleNumber(nums []int) int {
	n := len(nums)

	sort.Ints(nums)
	j := 0
	for i := 1; i < n; i += 2 {
		if nums[i] == nums[j] {
			j += 2
			continue
		}
		return nums[j]
	}

	return nums[n-1]
}

func singleNumberOfficial(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	sort.Ints(nums1)
	sort.Ints(nums2)
	j := 0
	k := 0
	if n1 >= n2 {
		for i := 0; i < n1; {
			if j >= n2 {
				break
			}
			for j < n2 {
				if i >= n1 {
					break
				}
				if nums1[i] == nums2[j] {
					nums2[k] = nums2[j]
					k++
					i++
					j++
					continue
				}
				if nums1[i] > nums2[j] {
					j++
				} else {
					i++
				}
			}
		}

		return nums2[:k]

	} else {
		for i := 0; i < n2; {
			if j >= n1 {
				break
			}
			for j < n1 {
				if i >= n2 {
					break
				}
				if nums2[i] == nums1[j] {
					nums1[k] = nums1[j]
					k++
					i++
					j++
					continue
				}
				if nums2[i] > nums1[j] {
					j++
				} else {
					i++
				}
			}
		}

		return nums1[:k]
	}
}

func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] > 8 {
		if n < 2 {
			return []int{1, 0}
		}
		return append(plusOne(digits[:n-1]), 0)
	} else {
		return append(digits[:n-1], digits[n-1]+1)
	}
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		v2, prs := mp[v]
		if prs {
			return []int{v2, i}
		}
		mp[target-v] = i
	}
	return nil
}

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		uniq := true
		mp := make(map[byte]bool)
		for _, v := range board[i] {
			if v == '.' {
				continue
			}
			_, prs := mp[v]
			if prs {
				uniq = false
				break
			}
			mp[v] = true
		}
		if !uniq {
			return false
		}
	}

	for j := range board {
		uniq := true
		mp := make(map[byte]bool)
		for i := range board {
			v := board[i][j]
			if v == '.' {
				continue
			}
			_, prs := mp[v]
			if prs {
				uniq = false
				break
			}
			mp[v] = true
		}
		if !uniq {
			return false
		}
	}

	for l := 0; l < 3; l++ {
		for k := 0; k < 3; k++ {
			uniq := true
			mp := make(map[byte]bool)
			for i := l * 3; i < l*3+3; i++ {
				for j := k * 3; j < k*3+3; j++ {
					v := board[i][j]
					if v == '.' {
						continue
					}
					_, prs := mp[v]
					if prs {
						uniq = false
						break
					}
					mp[v] = true
				}
			}
			if !uniq {
				return false
			}
		}
	}

	return true
}

func rotateImg(matrix [][]int) {
	N := len(matrix)
	n := N / 2
	m := N / 2
	if N%2 == 1 {
		m++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j], matrix[j][N-1-i], matrix[N-1-i][N-j-1], matrix[N-j-1][i] = matrix[N-j-1][i], matrix[i][j], matrix[j][N-1-i], matrix[N-1-i][N-j-1]
		}
	}
}
