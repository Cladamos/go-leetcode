package solutions

import (
	"slices"
	"testing"
)

func createList(nums []int) *ListNode {
	list := &ListNode{}
	prev := list
	for _, num := range nums {
		prev.Next = &ListNode{Val: num}
		prev = prev.Next
	}
	return list.Next
}

func createArray(list *ListNode) []int {
	var arr []int
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		val   int
		want  *ListNode
	}{
		{name: "Example 1", input: createList([]int{1, 2, 6, 3, 4, 5, 6}), val: 6, want: createList([]int{1, 2, 3, 4, 5})},
		{name: "Example 2", input: createList([]int{}), val: 1, want: createList([]int{})},
		{name: "Example 3", input: createList([]int{7, 7, 7, 7}), val: 7, want: createList([]int{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.input, tt.val)
			if !slices.Equal(createArray(got), createArray(tt.want)) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
