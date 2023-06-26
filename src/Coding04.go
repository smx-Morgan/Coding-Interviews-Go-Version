package src

func searchMatrix(matrix [][]int, target int) bool {
	rlen := len(matrix)
	if rlen <= 0 {
		return false
	}
	clen := len(matrix[0])
	if clen <= 0 {
		return false
	}
	index := clen - 1
	for c := 0; c < clen; c++ {
		if matrix[0][c] == target {
			return true
		}
		if matrix[0][c] > target {
			index = c - 1
		}
	}
	for r := 1; r < rlen; r++ {
		for x := 0; x <= index; x++ {
			if matrix[r][x] == target {
				return true
			}
		}
	}
	return false
}
