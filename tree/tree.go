package tree

import (
	"container/list"
	"fmt"
	"sort"
)

const (
	bigInt = 1<<31 + 32746
)

var pin = 0

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Queue struct {
	*list.List
}

func (q *Queue) in(v int) {
	q.PushBack(v)
}
func (q *Queue) out() int {
	v := q.Remove(q.Front()).(int)
	return v
}

func Main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil && root.Right != nil {
		return maxDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return maxDepth(root.Left) + 1
	}
	a, b := maxDepth(root.Left)+1, maxDepth(root.Right)+1
	if a >= b {
		return a
	} else {
		return b
	}
}

func minVal(tn *TreeNode) int {
	if tn == nil {
		return bigInt
	}
	maxLeft := minVal(tn.Left)
	maxRight := minVal(tn.Right)
	rank := []int{tn.Val, maxLeft, maxRight}
	sort.Ints(rank)
	return rank[0]
}
func maxVal(tn *TreeNode) int {
	if tn == nil {
		return -bigInt
	}
	maxLeft := maxVal(tn.Left)
	maxRight := maxVal(tn.Right)
	rank := []int{tn.Val, maxLeft, maxRight}
	sort.Ints(rank)
	return rank[2]
}

// 土法验证啊，运行时间击败0.46%的Go代码[哭]
func isValidBST(root *TreeNode) bool {
	fmt.Println(pin)
	pin++
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	fmt.Println(root.Val, minVal(root.Right))
	if root.Left == nil {
		return root.Val < minVal(root.Right) && isValidBST(root.Right)
	}
	if root.Right == nil {
		return maxVal(root.Left) < root.Val && isValidBST(root.Left)
	}
	return maxVal(root.Left) < root.Val && root.Val < minVal(root.Right) && isValidBST(root.Left) && isValidBST(root.Right)
}

var prev = -123454321
var valid = true

func isValidBSTMidOrder(root *TreeNode) bool {
	if root.Left != nil {
		isValidBSTMidOrder(root.Left)
	}
	fmt.Println(prev, root.Val)
	if prev >= root.Val {
		valid = false
	} else {
		prev = root.Val
	}
	if root.Right != nil {
		isValidBSTMidOrder(root.Right)
	}

	return valid
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if !(left != nil && right != nil) {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}

	return dfs(root.Left, root.Right)
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, *[][]int, int)
	dfs = func(tn *TreeNode, lo *[][]int, depth int) {
		if tn == nil {
			return
		}
		(*lo)[depth] = append((*lo)[depth], tn.Val)
		if !(tn.Left == nil && tn.Right == nil) && depth == len(*lo)-1 {
			fmt.Println(tn.Left, tn.Right)
			(*lo) = append(*lo, []int{})
		}
		dfs(tn.Left, lo, depth+1)
		dfs(tn.Right, lo, depth+1)
	}
	if root == nil {
		return res
	}
	res = append(res, []int{})
	dfs(root, &res, 0)
	return res
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	idx := n / 2
	root := &TreeNode{Val: nums[idx]}
	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}
