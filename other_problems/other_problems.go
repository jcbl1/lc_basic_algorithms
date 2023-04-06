package other_problems

import (
	"fmt"
	"sort"
)

func Main() {
	fmt.Println(isValid("]"))
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num % 2)
		num /= 2
	}
	return count
}

func hammingDistance(x int, y int) int {
	return hammingWeight(uint32(x ^ y))
}

func reverseBits(num uint32) uint32 {
	in := make([]bool, 32)
	for i := 0; num > 0; i++ {
		in[i] = num%2 == 1
		num /= 2
	}
	started := false
	var out uint32
	for i := 0; i < 32; i++ {
		if !started {
			if in[i] {
				started = true
				i--
			}
		} else {
			out *= 2
			if in[i] {
				out++
			}
		}
	}

	return out
}

type myStack struct {
	es []byte
}

func newStack() *myStack {
	return &myStack{}
}
func (st *myStack) push(v byte) {
	st.es = append(st.es, v)
}
func (st *myStack) pop() {
	st.es = st.es[:len(st.es)-1]
}
func (st myStack) top() byte {
	if len(st.es) < 1 {
		return 0
	}
	return st.es[len(st.es)-1]
}

func isValid(s string) bool {
	pStack := newStack()
	for _, v_ := range s {
		v := byte(v_)
		switch v {
		case "("[0]:
			pStack.push(v)
		case "["[0]:
			pStack.push(v)
		case "{"[0]:
			pStack.push(v)
		case ")"[0]:
			if pStack.top() != "("[0] {
				return false
			} else {
				pStack.pop()
			}
		case "]"[0]:
			if pStack.top() != "["[0] {
				return false
			} else {
				pStack.pop()
			}
		case "}"[0]:
			if pStack.top() != "{"[0] {
				return false
			} else {
				pStack.pop()
			}
		}
	}
	if len(pStack.es) != 0 {
		return false
	}
	return true
}

func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}
