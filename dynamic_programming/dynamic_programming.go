package dynamic_programming

import (
	"fmt"
	"math"
)

func Main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	for i := 2; i < n; i++ {
		tmp := n1 + n2
		n1, n2 = n2, tmp
	}
	return n2
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	minPrice := 1e9
	res := 0.
	for i := 1; i < n; i++ {
		minPrice = math.Min(minPrice, float64(prices[i-1]))
		res = math.Max(res, float64(prices[i]-int(minPrice)))
	}
	return int(res)
}

func sum(ns []int) int {
	res := 0
	for _, v := range ns {
		res += v
	}
	return res
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	realMax := nums[0]
	for j := 1; j < n; j++ {
		if nums[j] < nums[j-1]+nums[j] {
			nums[j] += nums[j-1]
		}
		if nums[j] > realMax {
			realMax = nums[j]
		}
	}
	return realMax
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	max := nums[1]
	if nums[0] > nums[1] {
		nums[1] = nums[0]
		max = nums[0]
	}
	for j := 2; j < n; j++ {
		if nums[j]+nums[j-2] > nums[j-1] {
			nums[j] += nums[j-2]
		} else {
			nums[j] = nums[j-1]
		}
		if nums[j] > max {
			max = nums[j]
		}
	}

	return max
}
