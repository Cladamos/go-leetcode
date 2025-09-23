package solutions

func isValidSudoku(board [][]byte) bool {

	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	boxes := make(map[int][]byte)

	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := board[row][col] - '1'
				boxKey := ((row / 3) * 3) + (col / 3)
				if rows[row][num] || cols[col][num] {
					return false
				}
				for _, i := range boxes[boxKey] {
					if i == num {
						return false
					}
				}
				rows[row][num] = true
				cols[col][num] = true
				boxes[boxKey] = append(boxes[boxKey], num)
			}

		}
	}

	return true
}
