package main

func CheckMovement(matrix [][]int) bool {
	lastPos := map[int]int{}
	robotIndex := 0

	for row := 0; row < len(matrix); row += 1 {
		robotIndex = 0

		for col := 0; col < len(matrix[row]); col += 1 {
			if matrix[row][col] == 1 {

				if row == 0 {
					lastPos[robotIndex] = col
				} else {
					if lastPos[robotIndex] >= col-1 && lastPos[robotIndex] <= col+1 {
						lastPos[robotIndex] = col
					} else if row != 0 {
						return false
					}
				}
				robotIndex += 1
			}

		}
	}

	return true
}
