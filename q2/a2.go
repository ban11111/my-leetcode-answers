package q2

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ten := 0
	preNode := &ListNode{}
	var next *ListNode = preNode
	for l1 != nil || l2 != nil {
		if l1 != nil {
			ten += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			ten += l2.Val
			l2 = l2.Next
		}
		next.Next = &ListNode{Val: ten % 10}
		next = next.Next
		ten /= 10
	}
	if ten > 0 {
		next.Next = &ListNode{Val: ten}
	}
	return preNode.Next
}
