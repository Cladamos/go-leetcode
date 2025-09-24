package solution

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name      string
		heightArr []int
		want      int
	}{
		{name: "Example 1", heightArr: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "Example 2", heightArr: []int{1, 1}, want: 1},
		{name: "Example 1", heightArr: []int{8, 7, 2, 1}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.heightArr)
			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
