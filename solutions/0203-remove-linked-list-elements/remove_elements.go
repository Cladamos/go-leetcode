package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tempList := &ListNode{Val: 0, Next: head}
	prev := tempList

	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	return tempList.Next
}
