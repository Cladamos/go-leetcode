package solutions

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{name: "Example 1", input: 121, want: true},
		{name: "Example 2", input: -121, want: false},
		{name: "Example 3", input: 10, want: false},
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
