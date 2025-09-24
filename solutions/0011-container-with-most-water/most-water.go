package solution

func maxArea(height []int) int {
	max := 0
	len := len(height) - 1

	l := 0
	r := len
	for l < r {

		lBar := height[l]
		rBar := height[r]
		small := lBar

		if lBar >= rBar {
			small = rBar
		}

		area := (r - l) * small
		if area > max {
			max = area
		}

		if lBar >= rBar {
			r--
		} else {
			l++
		}
	}
	return max
}
