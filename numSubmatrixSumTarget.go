package leetcode

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	rows := len(matrix) + 1
	cols := 0
	if rows != 0 {
		cols = len(matrix[0]) + 1
	}
	s := make([][]int, rows)
	s[0] = make([]int, cols)
	for i := 1; i < rows; i++ {
		s[i] = make([]int, cols)
		for j := 1; j < cols; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	ans := 0
	for x1 := 1; x1 < rows; x1++ {
		for y1 := 1; y1 < cols; y1++ {
			for x2 := 1; x2 <= x1; x2++ {
				for y2 := 1; y2 <= y1; y2++ {
					if s[x1][y1]-s[x2-1][y1]-s[x1][y2-1]+s[x2-1][y2-1] == target {
						ans += 1
					}
				}
			}
		}
	}
	return ans
}
