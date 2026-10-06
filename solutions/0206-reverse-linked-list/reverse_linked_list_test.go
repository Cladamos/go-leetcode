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

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{name: "Example 1", list: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "Example 2", list: []int{1, 2}, want: []int{2, 1}},
		{name: "Example 3", list: []int{}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(createList(tt.list))
			if !slices.Equal(createArray(got), tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
