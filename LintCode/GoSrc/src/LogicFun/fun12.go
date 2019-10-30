// keep studying -- 2019.10.30
package LogicFun

// 54. Spiral Matrix
func SpiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])
	res := make([]int, row*col)
	n := 0
	for start := 0; start < (row+1)/2 && start < (col+1)/2; start++ {
		i := start
		j := start
		if i == col-start-1 && j == col-start-1 {
			res[n] = matrix[i][j]
			continue
		}

		for ; j < col-start-1; j++ {
			res[n] = matrix[i][j]
			n++
		}
		for ; i < row-start-1; i++ {
			res[n] = matrix[i][j]
			n++
		}
		for ; j > start; j-- {
			res[n] = matrix[i][j]
			n++
		}
		for ; i > start; i-- {
			res[n] = matrix[i][j]
			n++
		}

	}
	return res
}
