// keep studying(i will go to company for my dream, the last day of the year) -- 2019.12.30
package LogicFun

import (
	"Common"
)

// 124. Binary Tree Maximum Path Sum
func MaxPathSum(root *Common.TreeNode) int {
	a, b := maxPathSumSub(root)
	if a > b {
		return a
	}
	return b
}

func maxPathSumSub(root *Common.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftVal := Common.MININTNUM
	rightVal := Common.MININTNUM
	leftAllVal := Common.MININTNUM
	rightAllVal := Common.MININTNUM
	if root.Left != nil {
		leftVal, leftAllVal = maxPathSumSub(root.Left)
	}
	if root.Right != nil {
		rightVal, rightAllVal = maxPathSumSub(root.Right)
	}

	firstVal := root.Val
	if rightVal > leftVal && rightVal > 0 {
		firstVal += rightVal
	} else if leftVal > rightVal && leftVal > 0 {
		firstVal += leftVal
	}

	secondVal := root.Val
	if rightVal > 0 {
		secondVal += rightVal
	}
	if leftVal > 0 {
		secondVal += leftVal
	}
	if rightAllVal > secondVal {
		secondVal = rightAllVal
	}
	if leftAllVal > secondVal {
		secondVal = leftAllVal
	}
	return firstVal, secondVal
}

// 125. Valid Palindrome
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	start := 0
	end := len(s) - 1
	startOK := false
	endOK := false
	for start < end {
		if startOK || ((s[start] >= 'a' && s[start] <= 'z') || (s[start] >= 'A' && s[start] <= 'Z') || (s[start] >= '0' && s[start] <= '9')) {
			startOK = true
		} else {
			start++
			continue
		}
		if endOK || ((s[end] >= 'a' && s[end] <= 'z') || (s[end] >= 'A' && s[end] <= 'Z') || (s[end] >= '0' && s[end] <= '9')) {
			endOK = true
		} else {
			end--
			continue
		}

		left := s[start]
		right := s[end]
		if left >= 'a' {
			left = left - 'a' + 'A'
		}
		if right >= 'a' {
			right = right - 'a' + 'A'
		}
		if left == right {
			start++
			end--
			startOK = false
			endOK = false
		} else {
			return false
		}

	}
	return true
}

// 126. Word Ladder II
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string
	if len(wordList) < 1 {
		return res
	}

	n := len(beginWord)
	wordMap := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return res
	}
	delete(wordMap, endWord)
	beginMap := map[string][][]string{beginWord: {{beginWord}}}
	beginTmpMap := make(map[string][][]string)
	endMap := map[string][][]string{endWord: {{endWord}}}
	endTmpMap := make(map[string][][]string)

	iscontinue := true
	for iscontinue {
		for key, ladders := range beginMap {
			for i := 0; i < n; i++ {
				for j := 0; j < Common.SMALL_ENGLISH_CHAR_NUM; j++ {
					newKey := []byte(key)
					if newKey[i] != 'a'+byte(j) {
						newKey[i] = 'a' + byte(j)
						newStr := string(newKey)
						if newStr == "ian" {
							println("haha")
						}
						endLadders, ok := endMap[newStr]
						if ok {
							iscontinue = false
							for _, ladder := range ladders {
								for _, endLadder := range endLadders {
									res = append(res, append(ladder, endLadder...))
								}
							}
						} else {
							_, ok = wordMap[newStr]
							if ok {
								newLadders := make([][]string, len(ladders))
								// deep-copy
								for i, _ := range newLadders {
									newLadders[i] = make([]string, len(ladders[i]))
									copy(newLadders[i], ladders[i])
									newLadders[i] = append(newLadders[i], newStr)
								}
								if _, ok = beginTmpMap[newStr]; ok {
									beginTmpMap[newStr] = append(beginTmpMap[newStr], newLadders...)
								} else {
									beginTmpMap[newStr] = newLadders
								}
							}
						}
					}
				}
			}
		}
		if !iscontinue {
			break
		}
		if len(beginTmpMap) == 0 {
			return res
		}
		for str, _ := range beginTmpMap {
			delete(wordMap, str)
		}
		beginMap = beginTmpMap
		beginTmpMap = make(map[string][][]string)

		// bottom -> top
		for key, ladders := range endMap {
			for i := 0; i < n; i++ {
				for j := 0; j < Common.SMALL_ENGLISH_CHAR_NUM; j++ {
					newKey := []byte(key)
					if newKey[i] != 'a'+byte(j) {
						newKey[i] = 'a' + byte(j)
						newStr := string(newKey)
						startLadders, ok := beginMap[newStr]
						if ok {
							iscontinue = false
							for _, ladder := range ladders {
								for _, startLadder := range startLadders {
									res = append(res, append(startLadder, ladder...))
								}
							}
						} else {
							_, ok = wordMap[newStr]
							if ok {
								newLadders := make([][]string, len(ladders))
								for i, _ := range newLadders {
									newLadders[i] = make([]string, len(ladders[i]))
									copy(newLadders[i], ladders[i])
									newLadders[i] = append([]string{newStr}, newLadders[i]...)
								}
								if _, ok = endTmpMap[newStr]; ok {
									endTmpMap[newStr] = append(endTmpMap[newStr], newLadders...)
								} else {
									endTmpMap[newStr] = newLadders
								}
							}
						}
					}
				}
			}
		}
		if len(endTmpMap) == 0 {
			return res
		}
		for str, _ := range endTmpMap {
			delete(wordMap, str)
		}
		endMap = endTmpMap
		endTmpMap = make(map[string][][]string)
	}
	return res
}

// 127. Word Ladder
func LadderLength(beginWord string, endWord string, wordList []string) int {
	res := 0
	if len(wordList) < 1 {
		return res
	}

	n := len(beginWord)
	wordMap := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return res
	}
	delete(wordMap, endWord)
	beginMap := map[string]struct{}{beginWord: struct{}{}}
	beginTmpMap := make(map[string]struct{})
	endMap := map[string]struct{}{endWord: struct{}{}}
	endTmpMap := make(map[string]struct{})

	for true {
		for key, _ := range beginMap {
			for i := 0; i < n; i++ {
				for j := 0; j < Common.SMALL_ENGLISH_CHAR_NUM; j++ {
					newKey := []byte(key)
					if newKey[i] != 'a'+byte(j) {
						newKey[i] = 'a' + byte(j)
						newStr := string(newKey)
						_, ok := endMap[newStr]
						if ok {
							return res + 2
						} else {
							_, ok = wordMap[newStr]
							if ok {
								beginTmpMap[newStr] = struct{}{}
							}
						}
					}
				}
			}
		}
		if len(beginTmpMap) == 0 {
			return 0
		}
		for str, _ := range beginTmpMap {
			delete(wordMap, str)
		}
		res++
		beginMap = beginTmpMap
		beginTmpMap = make(map[string]struct{})

		// bottom -> top
		for key, _ := range endMap {
			for i := 0; i < n; i++ {
				for j := 0; j < Common.SMALL_ENGLISH_CHAR_NUM; j++ {
					newKey := []byte(key)
					if newKey[i] != 'a'+byte(j) {
						newKey[i] = 'a' + byte(j)
						newStr := string(newKey)
						_, ok := beginMap[newStr]
						if ok {
							return res + 2
						} else {
							_, ok = wordMap[newStr]
							if ok {
								endTmpMap[newStr] = struct{}{}
							}
						}
					}
				}
			}
		}
		if len(endTmpMap) == 0 {
			return 0
		}
		for str, _ := range endTmpMap {
			delete(wordMap, str)
		}
		res++
		endMap = endTmpMap
		endTmpMap = make(map[string]struct{})
	}
	return res
}

// 128. Longest Consecutive Sequence
func LongestConsecutiveV2(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	res := 1
	numMap := make(map[int]int, len(nums))
	for _, num := range nums {
		numMap[num] = 1
	}

	for num, _ := range numMap {
		val := num + 1
		_, ok := numMap[val]
		for ; ok; _, ok = numMap[val] {
			numMap[num] = numMap[val] + numMap[num]
			if numMap[num] > res {
				res = numMap[num]
			}
			delete(numMap, val)
			val++
		}

	}
	return res
}

// 129. Sum Root to Leaf Numbers
func SumNumbers(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	res := 0
	curNum := root.Val
	if root.Left != nil {
		sumNumbersSub(root.Left, curNum, &res)
	}
	if root.Right != nil {
		sumNumbersSub(root.Right, curNum, &res)
	}

	return res
}

func sumNumbersSub(root *Common.TreeNode, curNum int, res *int) {
	curNum = curNum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res = *res + curNum
		return
	}

	if root.Left != nil {
		sumNumbersSub(root.Left, curNum, res)
	}
	if root.Right != nil {
		sumNumbersSub(root.Right, curNum, res)
	}
}

// 130. Surrounded Regions
func Solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}

	for i := 1; i < len(board[0])-1; i++ {
		if board[0][i] == 'O' {
			SolveSub(board, 1, i)
		}
		if board[len(board)-1][i] == 'O' {
			SolveSub(board, len(board)-2, i)
		}
	}
	for i := 1; i < len(board)-1; i++ {
		if board[i][0] == 'O' {
			SolveSub(board, i, 1)
		}
		if board[i][len(board[0])-1] == 'O' {
			SolveSub(board, i, len(board[0])-2)
		}
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func SolveSub(board [][]byte, i, j int) {
	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'
	if i-1 > 0 && board[i-1][j] != '#' {
		SolveSub(board, i-1, j)
	}
	if j-1 > 0 && board[i][j-1] != '#' {
		SolveSub(board, i, j-1)
	}
	if i+1 < len(board)-1 && board[i+1][j] != '#' {
		SolveSub(board, i+1, j)
	}
	if j+1 < len(board[0])-1 && board[i][j+1] != '#' {
		SolveSub(board, i, j+1)
	}
}

// 131. Palindrome Partitioning
func PartitionV2(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}

	var curStr []string
	partitionDp := partitionDP(s)
	partitionSub(s, 0, &curStr, &res, partitionDp)
	return res
}

func partitionDP(s string) [][]bool {
	res := make([][]bool, len(s))
	for i, _ := range res {
		res[i] = make([]bool, len(s))
		res[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if j-i == 1 || res[i+1][j-1] {
				if s[i] == s[j] {
					res[i][j] = true
				} else {
					res[i][j] = false
				}
			} else {
				res[i][j] = false
			}
		}
	}

	return res
}

func partitionSub(s string, start int, curStr *[]string, res *[][]string, dp [][]bool) {
	if start == len(s) {
		curRes := make([]string, len(*curStr))
		copy(curRes, *curStr)
		*res = append(*res, curRes)
		return
	}

	for i := start; i < len(s); i++ {
		if dp[start][i] {
			*curStr = append(*curStr, s[start:i+1])
			partitionSub(s, i+1, curStr, res, dp)
			*curStr = (*curStr)[:len(*curStr)-1]
		}
	}
}

/* error-> not partition,is sum chars
func partitionSub(s string, start int, curStr *[]string, res *[][]string) {
	if start == len(s) {
		curRes := make([]string, len(*curStr))
		copy(curRes, *curStr)
		*res = append(*res, curRes)
		return
	}

	*curStr = append(*curStr, s[start:start+1])
	partitionSub(s, start+1, curStr, res)
	*curStr = (*curStr)[:len(*curStr)-1]
	charMap := map[byte]int{s[start]: 1}
	mid := -1

	for i := start + 1; i < len(s); i++ {
		if (i-start)%2 == 1 && mid == -1 {
			charMap[s[i]]--
			if charMap[s[i]] == 0 {
				delete(charMap, s[i])
			}
			mid = i
		} else {
			charMap[s[mid]]++
			charMap[s[i]]--
			if charMap[s[mid]] == 0 {
				delete(charMap, s[mid])
			}
			if charMap[s[i]] == 0 {
				delete(charMap, s[i])
			}

			if (i-start)%2 == 1 {
				mid++
			}
		}

		if len(charMap) == 0 {
			*curStr = append(*curStr, s[start:i+1])
			partitionSub(s, i+1, curStr, res)
			*curStr = (*curStr)[:len(*curStr)-1]
		}
	}
}
*/

// 132. Palindrome Partitioning II
func MinCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	partitionDp := partitionDP(s)
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}

	for i := 1; i < len(s); i++ {
		if partitionDp[0][i] {
			dp[i] = 0
			continue
		}

		for j := 0; j < i; j++ {
			if partitionDp[j+1][i] {
				if dp[i] > dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[len(s)-1]
}

// 133 to PY(errors, wait for fix the bug)
// 134. Gas Station
func CanCompleteCircuit(gas []int, cost []int) int {
	start := -1
	sum := 0

	for i := 0; i < len(gas); i++ {
		if sum+gas[i]-cost[i] < 0 {
			if gas[i]-cost[i] < 0 {
				sum = 0
				start = -1
			} else {
				sum = gas[i] - cost[i]
				start = i
			}
		} else {
			sum += gas[i] - cost[i]
			if start == -1 {
				start = i
			}
		}
	}

	if start == -1 {
		return -1
	}

	for i := 0; i < start; i++ {
		sum += gas[i] - cost[i]
		if sum < 0 {
			return -1
		}
	}
	return start
}

// 135. Candy
func Candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	candySet := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candySet[i] = candySet[i-1] + 1
		}
	}

	sum := candySet[len(ratings)-1] + 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candySet[i] < candySet[i+1]+1 {
				candySet[i] = candySet[i+1] + 1
			}
		}
		sum += candySet[i] + 1
	}

	return sum
}

// 136. Single Number
func singleNumber(nums []int) int {
	numMap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			delete(numMap, num)
		} else {
			numMap[num] = struct{}{}
		}
	}

	for num, _ := range numMap {
		return num
	}

	return -1
}

// why slowly by using map ??
func singleNumberOpt(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

// 137. Single Number II
func SingleNumberV2(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		if v, ok := numMap[num]; ok && v == 2 {
			delete(numMap, num)
		} else {
			numMap[num]++
		}
	}

	for num, _ := range numMap {
		return num
	}

	return -1
}

func SingleNumberV2Opt(nums []int) int {
	a := 0
	b := 0

	for _, num := range nums {
		tmp := a & num
		a = a ^ num
		b = b ^ tmp

		tmp = a & b
		a = a - tmp
		b = b - tmp
	}

	return a ^ b
}
