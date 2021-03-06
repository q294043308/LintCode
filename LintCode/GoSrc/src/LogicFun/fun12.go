// keep studying -- 2019.10.30
package logicfun

import (
	"Common"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

// 60. Permutation Sequence
func GetPermutation(n int, k int) string {
	arv := make([]int, n)
	dev := 1
	for i := 0; i < n; i++ {
		arv[i] = i + 1
		dev = dev * (i + 1)
	}

	var res []byte
	for i := 0; i < n; i++ {
		dev = dev / (n - i)
		rem := (k-1)%dev + 1
		mod := (k - 1) / dev
		v := arv[mod]
		res = append(res, byte(v)+'0')
		k = rem
		arv = append(arv[:mod], arv[mod+1:]...)
	}
	return string(res)
}

// 61. Rotate List
func RotateRight(head *Common.ListNode, k int) *Common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	for node := head; true; node = node.Next {
		n++
		if node.Next == nil {
			node.Next = head
			break
		}
	}

	k = n - (k % n)
	tmp := head
	for i := 1; i < k; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next
	tmp.Next = nil
	return res
}

// 62. Unique Paths
func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	dynamicPath := make([][]int, m-1)
	for i := 0; i < m-1; i++ {
		dynamicPath[i] = make([]int, n-1)
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if i == 0 && j == 0 {
				dynamicPath[i][j] = 2
			} else if i == 0 {
				dynamicPath[i][j] = 1 + dynamicPath[i][j-1]
			} else if j == 0 {
				dynamicPath[i][j] = 1 + dynamicPath[i-1][j]
			} else {
				dynamicPath[i][j] = dynamicPath[i-1][j] + dynamicPath[i][j-1]
			}
		}
	}

	return dynamicPath[m-2][n-2]
}

// 62. Unique Paths
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dynamicPath := make([][]int, m)
	for i := 0; i < m; i++ {
		dynamicPath[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if i == 0 && j == 0 {
				dynamicPath[i][j] = 1
			} else if i == 0 {
				dynamicPath[i][j] = dynamicPath[i][j-1]
			} else if j == 0 {
				dynamicPath[i][j] = dynamicPath[i-1][j]
			} else {
				dynamicPath[i][j] = dynamicPath[i-1][j] + dynamicPath[i][j-1]
			}
		}
	}

	return dynamicPath[m-1][n-1]
}

// 64. Minimum Path Sum
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	dynamicPath := make([][]int, m)
	for i := 0; i < m; i++ {
		dynamicPath[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dynamicPath[i][j] = grid[i][j]
			} else if i == 0 {
				dynamicPath[i][j] = dynamicPath[i][j-1] + grid[i][j]
			} else if j == 0 {
				dynamicPath[i][j] = dynamicPath[i-1][j] + grid[i][j]
			} else {
				if dynamicPath[i-1][j] < dynamicPath[i][j-1] {
					dynamicPath[i][j] = dynamicPath[i-1][j] + grid[i][j]
				} else {
					dynamicPath[i][j] = dynamicPath[i][j-1] + grid[i][j]
				}
			}
		}
	}

	return dynamicPath[m-1][n-1]
}

// 65. Valid Number (i'm failed)
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	isSign := false
	haveE := false
	havePoint := false

	data := []byte(s)
	for len(data) > 0 {
		if data[0] == ' ' {
			data = data[1:]
		} else {
			break
		}
	}
	for len(data)-1 >= 0 {
		if data[len(data)-1] == ' ' {
			data = data[:len(data)-1]
		} else {
			break
		}
	}

	if len(data) == 0 {
		return false
	}

	for i := 0; i < len(data); i++ {
		if data[i] <= '9' && data[i] >= '0' {
			continue
		} else if data[i] == 'e' {
			if haveE {
				return false
			} else if i == 0 || i == len(data)-1 {
				return false
			}
			isSign = false
			haveE = true
		} else if data[i] == '+' || data[i] == '-' {
			if isSign {
				return false
			}
			if i != 0 && data[i-1] != 'e' {
				return false
			}
			if i == len(data)-1 {
				return false
			}
			isSign = true
		} else if data[i] == '.' {
			if havePoint || haveE {
				return false
			}
			if i == len(data)-1 && i == 0 {
				return false
			}
			if i != len(data)-1 && (data[i+1] > '9' || data[i+1] < '0') {
				return false
			}
			if i != 0 && (data[i-1] > '9' || data[i-1] < '0') && data[i-1] != '+' && data[i-1] != '-' {
				return false
			}
			havePoint = true
		} else {
			return false
		}
	}
	return true
}

// 65. Valid Number
func IsNumberV2(s string) bool {
	matched, _ := regexp.Match(`^ *[+-]?(\d+|\d+\.\d+|\d+\.|\.\d+)(e[+-]?\d+)? *$`, []byte(s))
	return matched
}

// 66. Plus One
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

// 67. Add Binary
func AddBinary(a string, b string) string {
	isAdd := false
	i := len(a) - 1
	j := len(b) - 1
	var res []byte

	for i >= 0 && j >= 0 {
		curV := a[i] - '0' + b[j] - '0'
		if isAdd {
			curV = curV + 1
			isAdd = false
		}
		if curV >= '2'-'0' {
			curV = curV - 2
			isAdd = true
		}
		res = append([]byte{curV + '0'}, res...)
		i--
		j--
	}

	for i >= 0 {
		if isAdd {
			curV := a[i] + 1
			if curV >= '2' {
				curV = curV - 2
			} else {
				isAdd = false
			}
			res = append([]byte{curV}, res...)
		} else {
			res = append([]byte(a)[:i+1], res...)
			break
		}
		i--
	}

	for j >= 0 {
		if isAdd {
			curV := b[j] + 1
			if curV >= '2' {
				curV = curV - 2
			} else {
				isAdd = false
			}
			res = append([]byte{curV}, res...)
		} else {
			res = append([]byte(b)[:j+1], res...)
			break
		}
		j--
	}

	if isAdd {
		res = append([]byte{'1'}, res...)
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}

// 68. Text Justification
func FullJustify(words []string, maxWidth int) []string {
	var res []string
	curV := []string{words[0]}
	length := len(words[0])

	for i := 1; i <= len(words); i++ {
		if i == len(words) || length+len(words[i])+1 > maxWidth {
			if length == 0 {
				break
			}

			curStr := ""
			if i == len(words) {
				for _, str := range curV {
					curStr = curStr + str + " "
				}
				curStr = curStr[:len(curStr)-1]
			} else {
				moreWidth := maxWidth - length
				allSpace := 0
				leftSpace := 0
				if len(curV) != 1 {
					allSpace = moreWidth / (len(curV) - 1)
					leftSpace = moreWidth % (len(curV) - 1)
				}

				for index, str := range curV {
					if index != len(curV)-1 {
						num := 1 + allSpace
						if leftSpace > 0 {
							leftSpace--
							num++
						}
						spaces := make([]byte, num)
						for i, _ := range spaces {
							spaces[i] = ' '
						}
						curStr = curStr + str + string(spaces)
					} else {
						curStr = curStr + str
					}
				}
			}

			if len(curStr) < maxWidth {
				spaces := make([]byte, maxWidth-len(curStr))
				for i, _ := range spaces {
					spaces[i] = ' '
				}
				curStr = curStr + string(spaces)
			}
			res = append(res, curStr)
			if i != len(words) {
				curV = []string{words[i]}
				length = len(words[i])
			}

		} else {
			length = len(words[i]) + 1 + length
			curV = append(curV, words[i])
		}
	}

	return res
}

// 69. Sqrt(x)
func MySqrt(x int) int {
	// Tailand Method - 泰勒展开式，忘了，得重新学习高数了
	return int(math.Sqrt(float64(x)))
}

// 70. Climbing Stairs
func ClimbStairs(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 2 {
		return 2
	}

	start := 1
	end := 2
	for i := 3; i <= n; i++ {
		tmp := start + end
		start = end
		end = tmp
	}
	return end
}

// 71. Simplify Path
func SimplifyPath(path string) string {
	res := make([]string, 0)
	vs := strings.Split(path, "/")
	for _, s := range vs {
		if len(s) == 0 {
			continue
		}

		if s[0] == '.' {
			if len(s) == 1 {
				continue
			}
			if len(s) == 2 {
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
				continue
			}
		}
		res = append(res, s)
	}
	return "/" + strings.Join(res, "/")
}

// 72. Edit Distance
func MinDistanceV2(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	n := len(word1)
	m := len(word2)
	disMap := make([][]int, n)
	for i := 0; i < n; i++ {
		disMap[i] = make([]int, m)
	}

	// first column
	isUse := false
	for i := 0; i < n; i++ {
		disMap[i][0] = i + 1
		if word1[i] == word2[0] {
			isUse = true
		}
		if isUse {
			disMap[i][0]--
		}
	}

	// first row
	isUse = false
	for i := 0; i < m; i++ {
		disMap[0][i] = i + 1
		if word1[0] == word2[i] {
			isUse = true
		}
		if isUse {
			disMap[0][i]--
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dis := Common.MAXINTNUM
			if word1[i] == word2[j] {
				if dis > disMap[i-1][j-1] {
					dis = disMap[i-1][j-1]
				}
			} else {
				if dis > disMap[i-1][j-1]+1 {
					dis = disMap[i-1][j-1] + 1
				}
			}
			if dis > disMap[i-1][j]+1 {
				dis = disMap[i-1][j] + 1
			}
			if dis > disMap[i][j-1]+1 {
				dis = disMap[i][j-1] + 1
			}
			disMap[i][j] = dis
		}
	}
	return disMap[n-1][m-1]
}

// 73. Set Matrix Zeroes
func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m := len(matrix)
	n := len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i, v := range rows {
		if v {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j, v := range cols {
		if v {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

// 74. Search a 2D Matrix
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n := len(matrix)
	m := len(matrix[0])
	start := 0
	end := n*m - 1

	for start <= end {
		mid := (end-start)/2 + start
		row := mid / m
		col := mid % m
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

// 75. Sort Colors
func SortColors(nums []int) {
	sort.Ints(nums)
}

// 76. Minimum Window Substring
func MinWindow(s string, t string) string {
	if len(s) == 0 {
		return ""
	}

	type node struct {
		index int
		char  byte
	}

	minstart := -1
	minend := Common.MAXINTNUM
	start := -1
	end := 0
	SourCharMap := make(map[byte][]int)
	SourIndexMap := make(map[int]struct{})
	charMap := make(map[byte]int)
	existMap := make(map[byte]struct{})
	for _, v := range t {
		charMap[byte(v)]++
		existMap[byte(v)] = struct{}{}
	}

	for end < len(s) {
		curV := s[end]
		if _, ok := existMap[curV]; ok {
			if start == -1 {
				start = end
			}

			num, ok := charMap[curV]
			if ok {
				if num == 1 {
					delete(charMap, curV)
				} else {
					charMap[curV]--
				}
			} else {
				if curV == s[start] {
					for true {
						start++
						oldV := s[start]
						if _, ok := existMap[oldV]; ok {
							if _, ok := SourIndexMap[start]; ok {
								SourCharMap[oldV] = SourCharMap[oldV][1:]
							} else if len(SourCharMap[oldV]) > 0 {
								delete(SourIndexMap, SourCharMap[oldV][len(SourCharMap[oldV])-1])
								SourCharMap[oldV] = SourCharMap[oldV][:len(SourCharMap[oldV])-1]

							} else {
								break
							}
						}
					}

				} else {
					SourCharMap[curV] = append(SourCharMap[curV], end)
					SourIndexMap[end] = struct{}{}
				}
			}

		}
		end++
		if len(charMap) == 0 && (minend-minstart) > (end-start) {
			minstart = start
			minend = end
		}
	}

	if minend > len(s) {
		return ""
	}
	return s[minstart:minend]
}

// 77. Combinations
func Combine(n int, k int) [][]int {
	var res [][]int
	if n == 0 || k == 0 {
		return res
	}
	cur := make([]int, 0, k)
	CombineSub(n, k, cur, &res, 1)
	return res
}

func CombineSub(n int, k int, cur []int, res *[][]int, start int) {
	if len(cur) == k {
		copyAr := make([]int, k)
		copy(copyAr, cur)
		*res = append(*res, copyAr)
		return
	}

	for start <= n-k+1+len(cur) {
		cur = append(cur, start)
		CombineSub(n, k, cur, res, start+1)
		cur = cur[:len(cur)-1]
		start++
	}
}

// 78. Subsets
func Subsets(nums []int) [][]int {
	res := [][]int{{}}
	if len(nums) == 0 {
		return res
	}
	var cur []int
	SubsetsSub(nums, cur, &res, 0)
	return res
}

func SubsetsSub(nums, cur []int, res *[][]int, index int) {
	if len(cur) != 0 {
		copyAr := make([]int, len(cur))
		copy(copyAr, cur)
		*res = append(*res, copyAr)
	}

	for index < len(nums) {
		cur = append(cur, nums[index])
		SubsetsSub(nums, cur, res, index+1)
		cur = cur[:len(cur)-1]
		index++
	}
}

// 79. Word Search
func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	signLoad := make([][]bool, len(board))
	for i := 0; i < len(signLoad); i++ {
		signLoad[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				signLoad[i][j] = true
				if isOk := ExistSub(board, signLoad, word[1:], i, j); isOk {
					return true
				}
				signLoad[i][j] = false
			}
		}
	}
	return false
}

func ExistSub(board [][]byte, signLoad [][]bool, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i > 0 && !signLoad[i-1][j] && word[0] == board[i-1][j] {
		signLoad[i-1][j] = true
		if isOk := ExistSub(board, signLoad, word[1:], i-1, j); isOk {
			return true
		}
		signLoad[i-1][j] = false
	}

	if i < len(board)-1 && !signLoad[i+1][j] && word[0] == board[i+1][j] {
		signLoad[i+1][j] = true
		if isOk := ExistSub(board, signLoad, word[1:], i+1, j); isOk {
			return true
		}
		signLoad[i+1][j] = false
	}

	if j > 0 && !signLoad[i][j-1] && word[0] == board[i][j-1] {
		signLoad[i][j-1] = true
		if isOk := ExistSub(board, signLoad, word[1:], i, j-1); isOk {
			return true
		}
		signLoad[i][j-1] = false
	}

	if j < len(board[0])-1 && !signLoad[i][j+1] && word[0] == board[i][j+1] {
		signLoad[i][j+1] = true
		if isOk := ExistSub(board, signLoad, word[1:], i, j+1); isOk {
			return true
		}
		signLoad[i][j+1] = false
	}
	return false
}

// 80. Remove Duplicates from Sorted Array II
func RemoveDuplicatesV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1
	lastNum := nums[0]
	curNum := 1
	for i := 1; i < len(nums); i++ {
		if lastNum != nums[i] {
			nums[j] = nums[i]
			j++
			lastNum = nums[i]
			curNum = 1
		} else {
			if curNum < 2 {
				curNum++
				nums[j] = nums[i]
				j++
			}
		}
	}

	return j
}

// 81. Search in Rotated Sorted Array II
func SearchV2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] == target {
			return true
		}

		for start != mid && nums[mid] == nums[start] {
			start++
			if start == mid {
				start = mid + 1
				goto OUT
			}
		}
		for end != mid && nums[mid] == nums[end] {
			end--
			if end == mid {
				end = mid - 1
				goto OUT
			}
		}

		if nums[mid] < nums[end] {
			if target > nums[mid] {
				if target == nums[end] {
					return true
				} else if target < nums[end] {
					start = mid + 1
					end--
				} else {
					end = mid - 1
				}
			} else {
				end = mid - 1
			}

		} else {
			if target < nums[mid] {
				if target == nums[start] {
					return true
				} else if target > nums[start] {
					start++
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				start = mid + 1
			}
		}
	OUT:
	}
	return false
}

// 82. Remove Duplicates from Sorted List II
func DeleteDuplicates(head *Common.ListNode) *Common.ListNode {
	if head == nil {
		return head
	}

	var lastNode *Common.ListNode
	curNode := head
	isEqual := false

	for curNode.Next != nil {
		if curNode.Val == curNode.Next.Val {
			curNode = curNode.Next
			isEqual = true
		} else {
			isEqual = false
			if lastNode == nil {
				if head == curNode {
					lastNode = head
					curNode = curNode.Next
				} else {
					head = curNode.Next
					curNode = head
				}
			} else {
				if lastNode.Next == curNode {
					lastNode = curNode
				} else {
					lastNode.Next = curNode.Next
				}
				curNode = curNode.Next
			}
		}
	}

	if isEqual {
		if lastNode == nil {
			return nil
		} else {
			lastNode.Next = nil
		}
	}

	return head
}

// 83. Remove Duplicates from Sorted List
func DeleteDuplicatesV2(head *Common.ListNode) *Common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curNode := head
	for curNode.Next != nil {
		if curNode.Val == curNode.Next.Val {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return head
}

// 84. Largest Rectangle in Histogram
func LargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	highs := make([]int, 0, len(heights))
	res := 0
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(highs); j++ {
			if highs[j] > heights[i] {
				highs[j] = heights[i]
			}
			curArx := highs[j] * (i - j + 1)
			if curArx > res {
				res = curArx
			}
		}
		if heights[i] > res {
			res = heights[i]
		}
		highs = append(highs, heights[i])
	}
	return res
}

// 84. Largest Rectangle in Histogram
func LargestRectangleAreaOpt(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := &Common.Stack{}
	res := 0

	type node struct {
		index int
		num   int
	}
	stack.Push(&node{index: -1, num: 0})
	for i := 0; i < len(heights); i++ {
		for stack.Size() > 1 {
			top := stack.Top().(*node)
			if heights[i] >= top.num {
				break
			}

			stack.Pop()
			nextTop := stack.Top().(*node)
			area := top.num * (i - nextTop.index - 1)
			if area > res {
				res = area
			}
		}
		stack.Push(&node{index: i, num: heights[i]})
	}

	for stack.Size() > 1 {
		top := stack.Pop().(*node)
		nextTop := stack.Top().(*node)
		area := top.num * (len(heights) - 1 - nextTop.index)
		if area > res {
			res = area
		}
	}
	return res
}

// 85. Maximal Rectangle
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	type node struct {
		height int
		left   int
		right  int
	}
	res := 0
	n := len(matrix)
	m := len(matrix[0])
	nodeMat := make([][]node, n)
	for i := 0; i < n; i++ {
		nodeMat[i] = make([]node, m)
	}

	continueNum := 0
	// full height and left
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				continueNum++
				if i == 0 {
					nodeMat[i][j].height = 1
				} else {
					nodeMat[i][j].height = 1 + nodeMat[i-1][j].height
				}

				if j == 0 {
					nodeMat[i][j].left = 1
				} else {
					nodeMat[i][j].left = continueNum
					if i != 0 && nodeMat[i-1][j].left != 0 && (nodeMat[i-1][j].left < nodeMat[i][j].left) {
						nodeMat[i][j].left = nodeMat[i-1][j].left
					}
				}
			} else {
				continueNum = 0
			}
		}
		continueNum = 0
	}

	// full right
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				continueNum++
				if j == m-1 {
					nodeMat[i][j].right = 1
				} else {
					nodeMat[i][j].right = continueNum
					if i != 0 && nodeMat[i-1][j].right != 0 && (nodeMat[i-1][j].right < nodeMat[i][j].right) {
						nodeMat[i][j].right = nodeMat[i-1][j].right
					}
				}
				curArx := nodeMat[i][j].height * (nodeMat[i][j].right + nodeMat[i][j].left - 1)
				if curArx > res {
					res = curArx
				}
			} else {
				continueNum = 0
			}
		}
		continueNum = 0
	}

	return res
}

// 86. Partition List
func Partition(head *Common.ListNode, x int) *Common.ListNode {
	if head == nil {
		return head
	}

	var newHead, newCurNode, lastNode, curNode *Common.ListNode
	if head.Val < x {
		newHead = head
		newCurNode = newHead
		head = head.Next
	}
	curNode = head

	for curNode != nil {
		if curNode.Val >= x {
			if lastNode == nil {
				lastNode = head
			} else {
				lastNode.Next = curNode
				lastNode = lastNode.Next
			}
		} else {
			if lastNode == nil {
				head = head.Next
			}
			if newHead == nil {
				newHead = curNode
				newCurNode = newHead
			} else {
				newCurNode.Next = curNode
				newCurNode = newCurNode.Next
			}
		}
		curNode = curNode.Next
	}

	if newHead != nil {
		newCurNode.Next = head
	} else {
		return head
	}

	if lastNode != nil {
		lastNode.Next = nil
	}

	return newHead
}

// 87. Scramble String
func IsScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	charMap := make(map[rune]int, Common.SMALL_ENGLISH_CHAR_NUM)
	for _, cha := range s1 {
		charMap[cha]++
	}
	for _, cha := range s2 {
		num, ok := charMap[cha]
		if !ok {
			break
		} else {
			if num == 1 {
				delete(charMap, cha)
			} else {
				charMap[cha]--
			}
		}
	}
	if len(charMap) != 0 {
		return false
	}

	for i := 1; i < len(s1); i++ {
		if IsScramble(s1[:i], s2[:i]) && IsScramble(s1[i:], s2[i:]) {
			return true
		}
		if IsScramble(s1[:i], s2[len(s1)-i:]) && IsScramble(s1[i:], s2[:len(s1)-i]) {
			return true
		}
	}
	return false
}

// 88. Merge Sorted Array
func MergeV2(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}

	nums1 = nums1[:m]
	i := 0
	j := 0
	for i < len(nums1) && j < n {
		if nums1[i] <= nums2[j] {
			i++
		} else {
			tmp := append([]int{nums2[j]}, nums1[i:]...)
			nums1 = append(nums1[:i], tmp...)
			i++
			j++
		}
	}

	if j < n {
		nums1 = append(nums1, nums2[j:]...)
	}
}

// 89. Gray Code
func GrayCode(n int) []int {
	res := []int{0}
	if n == 0 {
		return res
	}
	num := 0
	grayCodeSub(n-1, &num, &res)
	return res
}

func grayCodeSub(n int, num *int, res *[]int) {
	if n < 0 {
		return
	}

	grayCodeSub(n-1, num, res)
	*num = *num ^ (1 << uint(n))
	*res = append(*res, *num)
	grayCodeSub(n-1, num, res)
}

// 90. Subsets II

func SubsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		var curNode []int
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		subsetsWithDupSub(nums, &res, &curNode, i)
	}

	return res
}

func subsetsWithDupSub(nums []int, res *[][]int, curNode *[]int, index int) {
	if index >= len(nums) {
		return
	}

	*curNode = append(*curNode, nums[index])
	tmp := make([]int, len(*curNode))
	copy(tmp, *curNode)
	*res = append(*res, tmp)
	for i := index + 1; i < len(nums); {
		subsetsWithDupSub(nums, res, curNode, i)

		j := i + 1
		for j < len(nums) {
			if j < len(nums) {
				if (*curNode)[len(*curNode)-1] == nums[j] {
					j++
				} else {
					break
				}
			}
		}
		i = j
		*curNode = (*curNode)[:len(*curNode)-1]
	}
}

// 91. Decode Ways
func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	arr := make([]int, len(s))
	arr[0] = 1
	if s[1] == '0' && s[0] > '2' {
		return 0
	}
	if s[0] <= '2' {
		if (s[0] == '1' && s[1] != '0') || (s[1] <= '6' && s[1] > '0') {
			arr[1] = 2
		} else {
			arr[1] = 1
		}
	} else {
		arr[1] = 1
	}
	for i := 2; i < len(arr); i++ {
		if (s[i-1] > '2' || s[i-1] == '0') && s[i] == '0' {
			return 0
		}

		if s[i-1] > '2' || (s[i-1] == '2' && s[i] >= '7') || s[i-1] == '0' {
			arr[i] = arr[i-1]
		} else if s[i] == '0' {
			arr[i] = arr[i-2]
		} else {
			arr[i] = arr[i-2] + arr[i-1]
		}
	}

	return arr[len(arr)-1]
}

// 92. Reverse Linked List II
func ReverseBetween(head *Common.ListNode, m int, n int) *Common.ListNode {
	if m == n {
		return head
	}

	var pre, mNode *Common.ListNode
	curNode := head
	for curNode != nil {
		if m == 1 {
			mNode = pre
		}

		if n == 1 {
			// head
			if mNode == nil {
				head.Next = curNode.Next
				head = curNode
			} else {
				mNode.Next.Next = curNode.Next
				mNode.Next = curNode
			}
			curNode.Next = pre
			return head
		}

		if m <= 0 {
			next := curNode.Next
			curNode.Next = pre
			pre = curNode
			curNode = next
		} else {
			pre = curNode
			curNode = curNode.Next
		}
		m--
		n--
	}
	return head
}

// 93. Restore IP Addresses 25525511135
func RestoreIpAddresses(s string) []string {
	var res []string
	if len(s) < 4 {
		return res
	}

	var cur []int
	restoreIpAddressesSub(s, 0, &cur, &res)
	return res
}

func restoreIpAddressesSub(s string, start int, cur *[]int, res *[]string) {
	if len(*cur) == 4 {
		if start >= len(s) {
			curStr := make([]string, 4)
			for i, v := range *cur {
				curStr[i] = fmt.Sprintf("%d", v)
			}
			*res = append(*res, strings.Join(curStr, "."))
			return
		} else {
			return
		}
	}

	for i := 1; i <= 3 && start+i <= len(s); i++ {
		curS := s[start : start+i]
		curV, _ := strconv.Atoi(curS)
		if (curV == 0 && len(curS) > 1) || (curS[0] == '0' && curV > 0) || curV > 255 {
			return
		}

		*cur = append(*cur, curV)
		restoreIpAddressesSub(s, start+i, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}

// 94. Binary Tree Inorder Traversal
func InorderTraversal(root *Common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	if root.Left != nil {
		t := InorderTraversal(root.Left)
		res = append(res, t...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		t := InorderTraversal(root.Right)
		res = append(res, t...)
	}
	return res
}

// 95. Unique Binary Search Trees II
func GenerateTrees(n int) []*Common.TreeNode {
	if n == 0 {
		return nil
	}

	res := generateTreesSub(1, n)
	return res
}

func generateTreesSub(start, end int) []*Common.TreeNode {
	if start > end {
		return []*Common.TreeNode{nil}
	}

	var res []*Common.TreeNode
	for i := start; i <= end; i++ {
		left := generateTreesSub(start, i-1)
		right := generateTreesSub(i+1, end)
		for _, lv := range left {
			for _, rv := range right {
				curNode := &Common.TreeNode{
					Val: i,
				}
				curNode.Left = lv
				curNode.Right = rv
				res = append(res, curNode)
			}
		}

	}
	return res
}

// 96. Unique Binary Search Trees
func NumTrees(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, n)
	nums[0] = 1
	res := numTreesSub(1, n, &nums)
	return res
}

func numTreesSub(start, end int, nums *[]int) int {
	if start >= end {
		return 1
	}

	res := 0
	for i := start; i <= end; i++ {
		left := 0
		if i-1-start <= 0 {
			left = 1
		} else if (*nums)[i-1-start] != 0 {
			left = (*nums)[i-1-start]
		} else {
			left = numTreesSub(start, i-1, nums)
			(*nums)[i-1-start] = left
		}

		right := 0
		if end-i-1 <= 0 {
			right = 1
		} else if (*nums)[end-i-1] != 0 {
			right = (*nums)[end-i-1]
		} else {
			right = numTreesSub(i+1, end, nums)
			(*nums)[end-i-1] = right
		}

		res += left * right
	}
	return res
}

// 97. Interleaving String
func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 && len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if len(s1) != 0 && s1[0] == s3[0] {
		if IsInterleave(s1[1:], s2, s3[1:]) {
			return true
		}
	}
	if len(s2) != 0 && s2[0] == s3[0] {
		if IsInterleave(s1, s2[1:], s3[1:]) {
			return true
		}
	}
	return false
}

func IsInterleaveOpt(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 && len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// add more than col and row, for simple logic
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true

	for i := 0; i < len(s1); i++ {
		if s1[i] == s3[i] {
			dp[i+1][0] = true
		} else {
			break
		}
	}
	for j := 0; j < len(s2); j++ {
		if s2[j] == s3[j] {
			dp[0][j+1] = true
		} else {
			break
		}
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if dp[i][j+1] && s1[i] == s3[i+j+1] {
				dp[i+1][j+1] = true
				continue
			}
			if dp[i+1][j] && s2[j] == s3[i+j+1] {
				dp[i+1][j+1] = true
				continue
			}
			if dp[i][j] {
				if (s1[i] == s3[i+j] && s2[j] == s3[i+j+1]) || (s2[j] == s3[i+j] && s1[i] == s3[i+j+1]) {
					continue
				}
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// 98. Validate Binary Search Tree    ps:or using In-Order traversal(LMR)
func IsValidBST(root *Common.TreeNode) bool {
	_, _, ok := isValidBSTSub(root)
	return ok
}

func isValidBSTSub(root *Common.TreeNode) (int, int, bool) {
	if root == nil {
		return Common.MAXINTNUM, Common.MININTNUM, true
	}

	min := root.Val
	max := root.Val
	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return 0, 0, false
		}
		leftmin, leftmax, isOk := isValidBSTSub(root.Left)
		if leftmax >= root.Val || !isOk {
			return 0, 0, false
		}
		if leftmin < min {
			min = leftmin
		}
	}

	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return 0, 0, false
		}
		rightMin, rightMax, isOk := isValidBSTSub(root.Right)
		if rightMin <= root.Val || !isOk {
			return 0, 0, false
		}
		if rightMax > max {
			max = rightMax
		}
	}

	return min, max, true
}

// 99. Recover Binary Search Tree
func RecoverTree(root *Common.TreeNode) {
	var lastNode, firistNode, secondNode *Common.TreeNode
	recoverTreeSub(root, &lastNode, &firistNode, &secondNode)
	tmp := secondNode.Val
	secondNode.Val = firistNode.Val
	firistNode.Val = tmp
}

func recoverTreeSub(root *Common.TreeNode, lastNode, firistNode, secondNode **Common.TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		recoverTreeSub(root.Left, lastNode, firistNode, secondNode)
	}

	if *lastNode != nil {
		if *firistNode == nil && root.Val < (*lastNode).Val {
			(*firistNode) = *lastNode
		}
		if *firistNode != nil && root.Val < (*lastNode).Val {
			(*secondNode) = root
		}
	}
	*lastNode = root

	if root.Right != nil {
		recoverTreeSub(root.Right, lastNode, firistNode, secondNode)
	}
}

// 100. Same Tree
func IsSameTree(p *Common.TreeNode, q *Common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// 101. Symmetric Tree
func IsSymmetric(root *Common.TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricSub(root.Left, root.Right)

}

func isSymmetricSub(left, right *Common.TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return isSymmetricSub(left.Left, right.Right) && isSymmetricSub(left.Right, right.Left)
}

// 102. Binary Tree Level Order Traversal
func LevelOrder(root *Common.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var one, two []*Common.TreeNode
	one = append(one, root)
	for len(one)+len(two) > 0 {
		var curVal []int
		for _, node := range one {
			curVal = append(curVal, node.Val)
			if node.Left != nil {
				two = append(two, node.Left)
			}
			if node.Right != nil {
				two = append(two, node.Right)
			}
		}
		res = append(res, curVal)
		one = make([]*Common.TreeNode, 0)
		curVal = make([]int, 0)
		for _, node := range two {
			curVal = append(curVal, node.Val)
			if node.Left != nil {
				one = append(one, node.Left)
			}
			if node.Right != nil {
				one = append(one, node.Right)
			}
		}
		if len(curVal) > 0 {
			res = append(res, curVal)
		}
		two = make([]*Common.TreeNode, 0)
	}
	return res
}

// 103. Binary Tree Zigzag Level Order Traversal
func ZigzagLevelOrder(root *Common.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var one, two []*Common.TreeNode
	one = append(one, root)
	for len(one)+len(two) > 0 {
		var curVal []int
		for i := len(one) - 1; i >= 0; i-- {
			curNode := one[i]
			curVal = append(curVal, curNode.Val)
			if curNode.Left != nil {
				two = append(two, curNode.Left)
			}
			if curNode.Right != nil {
				two = append(two, curNode.Right)
			}
		}
		res = append(res, curVal)
		one = make([]*Common.TreeNode, 0)
		curVal = make([]int, 0)

		for i := len(two) - 1; i >= 0; i-- {
			curNode := two[i]
			curVal = append(curVal, curNode.Val)
			if curNode.Right != nil {
				one = append(one, curNode.Right)
			}
			if curNode.Left != nil {
				one = append(one, curNode.Left)
			}
		}
		if len(curVal) > 0 {
			res = append(res, curVal)
		}
		two = make([]*Common.TreeNode, 0)
	}
	return res
}

// 104. Maximum Depth of Binary Tree
func MaxDepth(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func BuildTree(preorder []int, inorder []int) *Common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &Common.TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = BuildTree(preorder[1:1+i], inorder[:i+1])
			root.Right = BuildTree(preorder[1+i:], inorder[i+1:])
			break
		}
	}

	return root
}

// 106. Construct Binary Tree from Inorder and Postorder Traversal
func BuildTreeV2(inorder []int, postorder []int) *Common.TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &Common.TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = BuildTreeV2(inorder[:i], postorder[:i])
			root.Right = BuildTreeV2(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}

	return root
}

// 107. Binary Tree Level Order Traversal II
func LevelOrderBottom(root *Common.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var one, two []*Common.TreeNode
	one = append(one, root)
	for len(one)+len(two) > 0 {
		var curVal []int
		for _, node := range one {
			curVal = append(curVal, node.Val)
			if node.Left != nil {
				two = append(two, node.Left)
			}
			if node.Right != nil {
				two = append(two, node.Right)
			}
		}
		res = append([][]int{curVal}, res...)
		one = make([]*Common.TreeNode, 0)
		curVal = make([]int, 0)
		for _, node := range two {
			curVal = append(curVal, node.Val)
			if node.Left != nil {
				one = append(one, node.Left)
			}
			if node.Right != nil {
				one = append(one, node.Right)
			}
		}
		if len(curVal) > 0 {
			res = append([][]int{curVal}, res...)
		}
		two = make([]*Common.TreeNode, 0)
	}
	return res
}

// 108. Convert Sorted Array to Binary Search Tree
func SortedArrayToBST(nums []int) *Common.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	ind := len(nums) / 2
	mid := nums[ind]
	root := &Common.TreeNode{
		Val:   mid,
		Left:  SortedArrayToBST(nums[:ind]),
		Right: SortedArrayToBST(nums[ind+1:]),
	}
	return root
}

// 109. Convert Sorted List to Binary Search Tree
func SortedListToBST(head *Common.ListNode) *Common.TreeNode {
	var nodeArr []int
	for head != nil {
		nodeArr = append(nodeArr, head.Val)
		head = head.Next
	}
	return sortedListToBSTSub(nodeArr)
}

func sortedListToBSTSub(nums []int) *Common.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	ind := len(nums) / 2
	mid := nums[ind]
	root := &Common.TreeNode{
		Val:   mid,
		Left:  sortedListToBSTSub(nums[:ind]),
		Right: sortedListToBSTSub(nums[ind+1:]),
	}
	return root
}

// 110. Balanced Binary Tree
func IsBalanced(root *Common.TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalancedSub(root) != -1
}

func isBalancedSub(root *Common.TreeNode) int {
	leftHight := 0
	rightHight := 0

	if root.Left != nil {
		leftHight = isBalancedSub(root.Left)
		if leftHight == -1 {
			return leftHight
		}
	}
	if root.Right != nil {
		rightHight = isBalancedSub(root.Right)
		if rightHight == -1 {
			return rightHight
		}
	}

	if leftHight > rightHight {
		if leftHight-rightHight > 1 {
			return -1
		} else {
			return leftHight + 1
		}
	} else {
		if rightHight-leftHight > 1 {
			return -1
		} else {
			return rightHight + 1
		}
	}
}

// 111. Minimum Depth of Binary Tree
func MinDepth(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := 0
	right := 0
	if root.Left != nil {
		left = MinDepth(root.Left)
	}
	if root.Right != nil {
		right = MinDepth(root.Right)
	}

	if left == 0 {
		return right + 1
	} else if right == 0 {
		return left + 1
	}

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

// 112. Path Sum
func HasPathSum(root *Common.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	if HasPathSum(root.Left, sum-root.Val) {
		return true
	}
	if HasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}

// 113. Path Sum II
func PathSum(root *Common.TreeNode, sum int) [][]int {
	var cur []int
	var res [][]int
	pathSumSub(root, sum, cur, &res)
	return res
}

func pathSumSub(root *Common.TreeNode, sum int, cur []int, res *[][]int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if sum == root.Val && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	pathSumSub(root.Left, sum-root.Val, cur, res)
	pathSumSub(root.Right, sum-root.Val, cur, res)
	return
}

// 114. Flatten Binary Tree to Linked List
func Flatten(root *Common.TreeNode) {
	var lastNode *Common.TreeNode
	flattenSub(root, &lastNode)
}

func flattenSub(root *Common.TreeNode, lastNode **Common.TreeNode) *Common.TreeNode {
	if root == nil {
		return root
	}

	right := root.Right
	*lastNode = root
	root.Right = flattenSub(root.Left, lastNode)
	if root.Left == nil {
		*lastNode = root
	} else {
		root.Left = nil
	}
	old := *lastNode
	old.Right = flattenSub(right, lastNode)
	return root
}

// 115. Distinct Subsequences
func NumDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}

	NumOp := make([][]int, len(t))
	for i := 0; i < len(NumOp); i++ {
		NumOp[i] = make([]int, len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i < len(t); i++ {
			base := 0
			if j != 0 {
				base = NumOp[i][j-1]
			}
			if s[j] == t[i] {
				if i == 0 || j == 0 {
					if i == 0 {
						NumOp[i][j] = base + 1
					} else {
						NumOp[i][j] = base
					}
				} else {
					NumOp[i][j] = base + NumOp[i-1][j-1]
				}
			} else {
				NumOp[i][j] = base
			}
		}
	}

	return NumOp[len(NumOp)-1][len(s)-1]
}

// 116.117 to C++
// 118. Pascal's Triangle
func Generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}

// 119. Pascal's Triangle II
func Generate2(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	res := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
	}

	for i := 2; i < rowIndex+1; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res[rowIndex]
}

// 120. Triangle
func MinimumTotal(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			base := triangle[i+1][j]
			if triangle[i+1][j+1] < base {
				base = triangle[i+1][j+1]
			}
			triangle[i][j] += base
		}
	}
	return triangle[0][0]
}

// 121. Best Time to Buy and Sell Stock
func MaxProfit(prices []int) int {
	res := 0
	if len(prices) == 0 || len(prices) == 1 {
		return res
	}

	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}

		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}

// 122. Best Time to Buy and Sell Stock II
func MaxProfitV2(prices []int) int {
	res := 0
	if len(prices) == 0 || len(prices) == 1 {
		return res
	}

	last := -1
	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) {
			if prices[i+1] > prices[i] {
				if last == -1 {
					last = prices[i]
				}
			} else if prices[i+1] < prices[i] {
				if last != -1 {
					res += prices[i] - last
					last = -1
				}
			}
		} else {
			if last != -1 {
				res += prices[i] - last
			}
		}
	}
	return res
}

// 122. Best Time to Buy and Sell Stock III
func MaxProfitV3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0

	pricesOp := make([][]int, 4)
	for i := 0; i < len(pricesOp); i++ {
		pricesOp[i] = make([]int, len(prices))
	}

	// buy
	for i := 0; i < len(prices); i++ {
		pricesOp[0][i] = -prices[i]
	}

	// sale
	maxnum := Common.MININTNUM
	for i := 1; i < len(prices); i++ {
		if pricesOp[0][i-1] > maxnum {
			maxnum = pricesOp[0][i-1]
		}
		pricesOp[1][i] = maxnum + prices[i]
		if res < pricesOp[1][i] {
			res = pricesOp[1][i]
		}
	}
	if len(prices) < 4 {
		return res
	}

	// buy
	maxnum = Common.MININTNUM
	for i := 2; i < len(prices); i++ {
		if pricesOp[1][i-1] > maxnum {
			maxnum = pricesOp[1][i-1]
		}
		pricesOp[2][i] = maxnum - prices[i]
	}

	// sale
	maxnum = Common.MININTNUM
	for i := 3; i < len(prices); i++ {
		if pricesOp[2][i-1] > maxnum {
			maxnum = pricesOp[2][i-1]
		}
		pricesOp[3][i] = maxnum + prices[i]
		if res < pricesOp[3][i] {
			res = pricesOp[3][i]
		}
	}

	return res
}
