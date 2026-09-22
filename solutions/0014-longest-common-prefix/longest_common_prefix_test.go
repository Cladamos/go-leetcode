package solutions

import "testing"

func TestIsLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "Example 1", input: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "Example 2", input: []string{"dog", "racecar", "car"}, want: ""},
		{name: "Example 3", input: []string{"a"}, want: "a"},
		{name: "Example 4", input: []string{""}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
