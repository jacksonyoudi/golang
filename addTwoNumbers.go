package main

import (
	"strconv"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var m []int
	p := l1
	for {
		if p.Next != nil {
			m = append(m, p.Val)
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
		if q.Next != nil {
			n = append(n, p.Val)
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
		fmt.Println(i)
		fmt.Println(s)
		strconv.Atoi(string(s))
		//l1 = ListNode{Val:varV, Next:nil}

	}
}

func main()  {
	
}
