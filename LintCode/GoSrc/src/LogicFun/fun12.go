// keep studying -- 2019.10.30
package LogicFun

// 54. Spiral Matrix
func SpiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}

	col := len(matrix[0])
	res := make([]int, row*col)
	n := 0
	for start := 0; start < (row+1)/2 && start < (col+1)/2; start++ {
		i := start
		j := start
		cj := false
		if i == row-start-1 && j == col-start-1 {
			res[n] = matrix[i][j]
			continue
		}

		if j == col-start-1 {
			cj = true
		}
		for ; j < col-start-1; j++ {
			res[n] = matrix[i][j]
			n++
		}
		if i == row-start-1 {
			res[n] = matrix[i][j]
			continue
		}
		for ; i < row-start-1; i++ {
			res[n] = matrix[i][j]
			n++
		}
		for ; j > start; j-- {
			res[n] = matrix[i][j]
			n++
		}
		if cj {
			res[n] = matrix[i][j]
			continue
		}
		for ; i > start; i-- {
			res[n] = matrix[i][j]
			n++
		}

	}
	return res
}

// 55. Jump Game
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	lastTrue := len(nums) - 1
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] >= lastTrue-i {
			lastTrue = i
		}
	}
	return nums[0] >= lastTrue-0
}

func MergeSub(arr [][]int, start, end int) {
	if start >= end {
		return
	}

	i := start
	j := end
	left := true
	for i < j {
		if left {
			if arr[i][0] > arr[j][0] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
				left = false
				j--
			} else {
				i++
			}
		} else {
			if arr[j][0] < arr[i][0] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
				left = true
				i++
			} else {
				j--
			}
		}
	}
	MergeSub(arr, start, i-1)
	MergeSub(arr, i+1, end)
}

// 56. Merge Intervals
func Merge(intervals [][]int) [][]int {
	var res [][]int
	MergeSub(intervals, 0, len(intervals)-1)
	for _, v := range intervals {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			if v[0] <= res[len(res)-1][1] {
				if v[1] > res[len(res)-1][1] {
					res[len(res)-1][1] = v[1]
				}
			} else {
				res = append(res, v)
			}
		}
	}
	return res
}

// 57. Insert Interval
func Insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	for _, v := range intervals {
		if v[0] > newInterval[1] || v[1] < newInterval[0] {
			res = append(res, v)
		} else {
			if v[0] < newInterval[0] {
				newInterval[0] = v[0]
			}
			if v[1] > newInterval[1] {
				newInterval[1] = v[1]
			}
		}
	}

	for i, v := range res {
		if v[0] > newInterval[0] {
			// tmp := make([][]int, 0, len(i+1))
			tmp := append([][]int{newInterval}, res[i:]...)
			res = append(res[:i], tmp...)
			return res
		}
	}
	res = append(res, newInterval)
	return res
}

// 58. Length of Last Word
func LengthOfLastWord(s string) int {
	res := 0
	start := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if start {
				return res
			}
		} else {
			start = true
			res++
		}
	}
	return res
}

// 59. Spiral Matrix II
func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}

	num := 1
	for start := 0; start < (n+1)/2; start++ {
		i := start
		j := start
		cj := false
		if i == n-start-1 && j == n-start-1 {
			res[i][j] = num
			continue
		}

		if j == n-start-1 {
			cj = true
		}
		for ; j < n-start-1; j++ {
			res[i][j] = num
			num++
		}
		if i == n-start-1 {
			res[i][j] = num
			continue
		}
		for ; i < n-start-1; i++ {
			res[i][j] = num
			num++
		}
		for ; j > start; j-- {
			res[i][j] = num
			num++
		}
		if cj {
			res[i][j] = num
			continue
		}
		for ; i > start; i-- {
			res[i][j] = num
			num++
		}

	}
	return res
}
