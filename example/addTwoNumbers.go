package main

import (
	"strconv"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	var m []int
	p := l1
	for {
		m = append(m, p.Val)
		if p.Next != nil {
			p = p.Next
		} else {
			break
		}
	}

	l1Ten := 0
	for i, v := range m {
		if i == 0 {
			l1Ten += v
		} else {
			l1Ten += i * 10 * v
		}
	}

	var n []int
	q := l2
	for {
		n = append(n, p.Val)
		if q.Next != nil {
			p = p.Next
		} else {
			break
		}
	}

	l2Ten := 0
	for i, v := range n {
		if i == 0 {
			l2Ten += v
		} else {
			l2Ten += i * 10 * v
		}
	}

	x := l1Ten + l2Ten
	y := strconv.Itoa(x)
	for i,s := range  []byte(y) {
		fmt.Println(i, s)

	}
}

func main()  {
	l1 := &ListNode{2, &ListNode{4, &ListNode{Val:3}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{Val:7}}}


	addTwoNumbers(l1, l2)
}
