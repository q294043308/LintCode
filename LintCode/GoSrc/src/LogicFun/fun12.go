// keep studying -- 2019.10.30
package LogicFun

import (
	"Common"
	"math"
	"regexp"
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
