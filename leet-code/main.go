package leetcode

import "fmt"

func Init() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	result := addTwoNumbers(l1, l2)
	fmt.Printf("result: %v\n", result)
}
