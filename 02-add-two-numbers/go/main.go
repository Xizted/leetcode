package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		a := 0
		b := 0

		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry

		carry = sum / 10
		result := sum % 10

		current.Next = &ListNode{
			Val: result,
		}
		current = current.Next

	}

	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyHead.Next
}

func main() {

	l1 := ListNode{
		Val:  5,
		Next: nil,
	}

	l2 := ListNode{
		Val:  5,
		Next: nil,
	}

	result := addTwoNumbers(&l1, &l2)

	fmt.Print(result)

}
