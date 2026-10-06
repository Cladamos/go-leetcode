package solutions

import "testing"

func createCycleList(nums []int, pos int) *ListNode {
	list := &ListNode{}
	prev := list

	lastIndex := len(nums) - 1
	for i, num := range nums {
		if i == lastIndex && pos >= 0 {
			cycleNode := list
			for pos > 0 {
				cycleNode = cycleNode.Next
				pos--
			}
			prev.Next = cycleNode
			break
		}
		prev.Next = &ListNode{Val: num}
		prev = prev.Next
	}
	return list.Next
}

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		want  bool
	}{
		{name: "Example 1", input: createCycleList([]int{3, 2, 0, -4}, 1), want: true},
		{name: "Example 2", input: createCycleList([]int{1, 2}, 0), want: true},
		{name: "Example 3", input: createCycleList([]int{1}, -1), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasCycle(tt.input)
			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
