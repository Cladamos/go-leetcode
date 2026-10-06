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

func TestDeleteNode(t *testing.T) {
	testList := createList([]int{4, 5, 1, 9})
	tests := []struct {
		name  string
		input *ListNode
		want  []int
	}{
		// Delete node2 in testList
		{name: "Example 1", input: testList.Next, want: []int{4, 1, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.input)
			got := createArray(testList)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
