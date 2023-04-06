package design_problem

import (
	"fmt"
	// "math/rand"
)

func Main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

// type Solution struct {
// 	Elems []int
// }

// func Constructor(nums []int) Solution {
// 	sol:=Solution{nums}
// 	return sol
// }

// func (this *Solution) Reset() []int {
// 	return this.Elems
// }

// func (this *Solution) Shuffle() []int {
// 	n:=len(this.Elems)
// 	origins:=make([]int,n)
// 	copy(origins,this.Elems)
// 	res:=make([]int,n)
// 	for i:=range res{
// 		randNum:=0
// 		if n-i-1>0{
// 			randNum=rand.Intn(n-i-1)
// 		}
// 		res[i]=origins[randNum]
// 		origins[randNum],origins[n-i-1]=origins[n-i-1],origins[randNum]
// 	}
// 	return res
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

type MinStack struct {
	mins   []int
	nums   []int
	height int
}

func Constructor() MinStack {
	return MinStack{nil, nil, 0}
}

func (this *MinStack) Push(val int) {
	if this.height == 0 {
		this.mins = append(this.mins, val)
		this.nums = append(this.nums, val)
		this.height++
		return
	}
	if this.mins[this.height-1] > val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[this.height-1])
	}
	this.nums = append(this.nums, val)
	this.height++
}

func (this *MinStack) Pop() {
	if this.height == 0 {
		return
	}
	fmt.Println("Poping: ", this.mins, this.nums, this.height)
	this.height--
	this.mins = this.mins[:this.height]
	this.nums = this.nums[:this.height]
}

func (this *MinStack) Top() int {
	return this.nums[this.height-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.height-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
