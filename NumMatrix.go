package leetcode

type NumMatrix struct {
	sum [][]int
}

func Constructor_1(matrix [][]int) NumMatrix {
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
	return NumMatrix{sum: s}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return m.sum[row2+1][col2+1] - m.sum[row1][col2+1] - m.sum[row2+1][col1] + m.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
