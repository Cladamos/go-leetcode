package solutions

import "slices"

func longestCommonPrefix(strs []string) string {
	slices.Sort(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	i := 0
	for i < len(first) && first[i] == last[i] {
		i++
	}
	return first[:i]
}
