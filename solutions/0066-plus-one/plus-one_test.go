package solution

import (
	"slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Example 1", input: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{name: "Example 2", input: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		{name: "Example 3", input: []int{9}, want: []int{1, 0}},
		{name: "Example 4", input: []int{9, 9}, want: []int{1, 0, 0}},
		{name: "Example 5", input: []int{9, 9, 9}, want: []int{1, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
