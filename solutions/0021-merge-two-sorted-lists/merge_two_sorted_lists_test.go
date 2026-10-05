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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{name: "Example 1", list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}},
		{name: "Example 2", list1: []int{}, list2: []int{}, want: []int{}},
		{name: "Example 3", list1: []int{}, list2: []int{0}, want: []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(createList(tt.list1), createList(tt.list2))
			if !slices.Equal(createArray(got), tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
