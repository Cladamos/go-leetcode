package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Revert second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	isPalindrome := true
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			isPalindrome = false
		}
		head = head.Next
		prev = prev.Next

	}
	return isPalindrome
}
