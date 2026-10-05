package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	prev := newList

	n1, n2 := list1, list2
	for n1 != nil && n2 != nil {
		if n1.Val > n2.Val {
			prev.Next = n2
			n2 = n2.Next
		} else {
			prev.Next = n1
			n1 = n1.Next
		}

		prev = prev.Next
	}
	if n1 == nil {
		prev.Next = n2
	} else {
		prev.Next = n1
	}

	return newList.Next
}
