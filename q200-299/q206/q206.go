package q206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListV1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var reversedHead *ListNode
	reversedTail := &ListNode{Val: head.Val, Next: nil}
	for next := head.Next; next != nil; {
		reversedHead = &ListNode{Val: next.Val, Next: reversedTail}
		reversedTail = reversedHead
		next = next.Next
	}
	return reversedHead
}

// 其实没必要重新创建一个链表，直接修改原链表的指针就行了, 按顺序反过来修改指针就行了
// 哈哈, 就像是把箭头一个一个掰过来调转方向
func reverseList(head *ListNode) *ListNode {
	var reversedHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = reversedHead
		reversedHead = head
		head = next
	}
	return reversedHead
}
