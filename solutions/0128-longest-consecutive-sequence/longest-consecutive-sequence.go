package solutions

func longestConsecutive(nums []int) int {

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	longestSequence := 0

	for num := range m {
		if !m[num-1] {
			curr := num
			len := 1

			for m[curr+1] {
				curr++
				len++
			}
			if len > longestSequence {
				longestSequence = len
			}
		}
	}
	return longestSequence
}
