package solutions

import "testing"

func TestLongestSequence(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		sequenceLen int
	}{
		{name: "Example 1", arr: []int{1, 2, 3, 4}, sequenceLen: 4},
		{name: "Example 2", arr: []int{}, sequenceLen: 0},
		{name: "Example 3", arr: []int{1, 7, 1, 8}, sequenceLen: 2},
		{name: "Example 4", arr: []int{1, 0}, sequenceLen: 2},
		{name: "Example 5", arr: []int{15, 16, 3, 4}, sequenceLen: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.arr)
			if tt.sequenceLen != got {
				t.Errorf("got %v, want %v", got, tt.sequenceLen)

			}
		})

	}
}
