// keep studying(i will go to company for my dream, the last day of the year) -- 2019.12.30
package LogicFun

import "Common"

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
