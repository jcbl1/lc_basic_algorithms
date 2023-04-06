package linked_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	elems := ln.assArr()
	return fmt.Sprint(elems)
}
func (ln *ListNode) assArr() []int {
	res := []int{ln.Val}
	if ln.Next == nil {
		return res
	}
	res = append(res, ln.Next.assArr()...)
	return res
}
func (ln *ListNode) Append(x ...int) {
	if len(x) == 0 {
		return
	}
	if ln.Next == nil {
		node := &ListNode{Val: x[0]}
		ln.Next = node
		ln.Next.Append(x[1:]...)
	} else {
		ln.Next.Append(x...)
	}
}

func Main() {
	head := &ListNode{Val: 4}
	head.Append(5)
	node := &ListNode{Val: 1}
	head.Next.Next = node
	node.Append(9)
	fmt.Println(head)
	deleteNode(node)
	fmt.Println(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func getSize(ln *ListNode) int {
	if ln.Next == nil {
		return 1
	}
	return 1 + getSize(ln.Next)
}
func nthNode(ln *ListNode, n int) *ListNode {
	if n == 0 {
		return ln
	}
	return nthNode(ln.Next, n-1)
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := getSize(head)
	n_ := sz - n
	node := nthNode(head, n_)
	deleteNode(node)
	return head
}
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	sz := getSize(head)
	cur := dummy
	for i := 0; i < n-sz-1; i++ {
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = cur.Next.Next
	return head
}
