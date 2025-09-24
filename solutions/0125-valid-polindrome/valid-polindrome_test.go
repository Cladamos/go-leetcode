package solution

import "testing"

func TestIsValidPolindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Example 1", input: "A man, a plan, a canal: Panama", want: true},
		{name: "Example 2", input: "race a car", want: false},
		{name: "Example 3", input: " ", want: true},
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
