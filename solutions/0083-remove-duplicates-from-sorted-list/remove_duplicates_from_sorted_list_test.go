package solution

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

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		want  *ListNode
	}{
		{name: "Example 1", input: createList([]int{1, 1, 2}), want: createList([]int{1, 2})},
		{name: "Example 2", input: createList([]int{1, 1, 2, 3, 3}), want: createList([]int{1, 2, 3})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.input)
			if !slices.Equal(createArray(got), createArray(tt.want)) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
