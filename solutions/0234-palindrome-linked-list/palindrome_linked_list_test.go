package solutions

import "testing"

func createList(nums []int) *ListNode {
	list := &ListNode{}
	prev := list
	for _, num := range nums {
		prev.Next = &ListNode{Val: num}
		prev = prev.Next
	}
	return list.Next
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		want  bool
	}{
		{name: "Example 1", input: createList([]int{1, 2, 2, 1}), want: true},
		{name: "Example 2", input: createList([]int{1, 2}), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
