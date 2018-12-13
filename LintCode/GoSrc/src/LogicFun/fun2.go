package LogicFun

import (
	"Common"
	"math"
	"sort"
	"strings"
)

// 1207. Teemo Attacking        Captain Teemo is on standby
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	sort.Ints(timeSeries)
	lastSerie := timeSeries[0]
	curSerie := timeSeries[0]
	res := 0

	for i := 1; i <= len(timeSeries); i++ {
		if i == len(timeSeries) || curSerie+duration < timeSeries[i] {
			res += curSerie - lastSerie + duration

			if i == len(timeSeries) {
				break
			}

			curSerie = timeSeries[i]
			lastSerie = timeSeries[i]
		} else {
			curSerie = timeSeries[i]
		}
	}
	return res
}

// 1394. Goat Latin
func ToGoatLatin(S string) string {
	if len(S) == 0 {
		return S
	}

	var res string
	appendStr := "maa"
	startIndex := 0
	vowelMap := make(map[byte]struct{})
	vowelMap['a'] = struct{}{}
	vowelMap['e'] = struct{}{}
	vowelMap['i'] = struct{}{}
	vowelMap['o'] = struct{}{}
	vowelMap['u'] = struct{}{}
	vowelMap['A'] = struct{}{}
	vowelMap['E'] = struct{}{}
	vowelMap['I'] = struct{}{}
	vowelMap['O'] = struct{}{}
	vowelMap['U'] = struct{}{}

	for {
		endIndex := strings.Index(S[startIndex:], " ") + startIndex
		var curStr string
		if endIndex < startIndex {
			curStr = S[startIndex:]
			startIndex = -1
		} else {
			curStr = S[startIndex:endIndex]
			startIndex = endIndex + 1
		}

		if len(curStr) == 0 {
			if startIndex == -1 {
				break
			}
			res += " "
			continue
		}

		if _, ok := vowelMap[curStr[0]]; !ok {
			curStr = curStr[1:] + string(curStr[0]) + appendStr
		} else {
			curStr += appendStr
		}
		appendStr += "a"

		res += curStr
		if startIndex != -1 {
			res += " "
		} else {
			break
		}
	}

	return res
}

// 1369. Most Common Word
func MostCommonWord(paragraph string, banned []string) string {
	strCountMap := make(map[string]int)

	for _, str := range banned {
		strCountMap[str] = -1
	}

	startIndex := 0
	var res string
	maxCount := 0

	for {
		endIndexPoint1 := strings.Index(paragraph[startIndex:], ",") + startIndex
		endIndexSpace := strings.Index(paragraph[startIndex:], " ") + startIndex
		endIndexPoint := strings.Index(paragraph[startIndex:], ".") + startIndex
		endIndexaa := strings.Index(paragraph[startIndex:], "!") + startIndex
		if endIndexPoint1 < startIndex {
			endIndexPoint1 = 100000
		}
		if endIndexSpace < startIndex {
			endIndexSpace = 100000
		}
		if endIndexPoint < startIndex {
			endIndexPoint = 100000
		}
		if endIndexaa < startIndex {
			endIndexaa = 100000
		}
		endIndex := int(math.Min(math.Min(math.Min(float64(endIndexPoint1), float64(endIndexSpace)), float64(endIndexPoint)), float64(endIndexaa)))
		var curStr string
		if endIndex < startIndex || endIndex == 100000 {
			curStr = paragraph[startIndex:]
			startIndex = -1
		} else {
			curStr = paragraph[startIndex:endIndex]
			startIndex = endIndex + 1
		}

		if len(curStr) == 0 {
			if startIndex == -1 {
				break
			}
			continue
		}

		curStr = strings.ToLower(curStr)
		if _, ok := strCountMap[curStr]; !ok {
			strCountMap[curStr] = 1
			if maxCount == 0 {
				maxCount = 1
				res = curStr
			}
		} else {
			if strCountMap[curStr] != -1 {
				strCountMap[curStr]++
				if strCountMap[curStr] > maxCount {
					maxCount = strCountMap[curStr]
					res = curStr
				}
			}
		}
		if startIndex == -1 {
			break
		}
	}
	return res
}

// 1201. Next Greater Element II
func NextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	numStack := new(Common.Stack)

	for i := 0; i < 2*len(nums); i++ {
		index := i % len(nums)
		topNum, _ := numStack.Top().(int)
		for !numStack.Empty() && nums[topNum] < nums[index] {
			popNum, _ := numStack.Pop().(int)
			res[popNum] = nums[index]
			topNum, _ = numStack.Top().(int)
		}
		if i < len(nums) {
			res[index] = -1
			numStack.Push(index)
		}
	}
	return res
}

// 1089. Valid Parenthesis String
func CheckValidString(s string) bool {
	left := 0  // 将所有*当做左括号时左括号剩余的个数
	right := 0 // 将所有*当做右括号或空时左括号剩余的个数

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
			right++
		} else {
			if right > 0 {
				right--
			}
			if s[i] == ')' {
				left--
			} else {
				left++
			}
		}
		if left < 0 {
			return false
		}
	}
	return right == 0
}

// 873. Squirrel Simulation
func MinDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	res := 0
	firstSelect := -1
	distanceSelect := float64(0)
	distanceToTreeSelect := float64(0)
	maxSelect := float64(Common.MININTNUM)

	for i := 0; i < len(nuts); i++ {
		toTree := math.Abs(float64(nuts[i][0]-tree[0])) + math.Abs(float64(nuts[i][1]-tree[1]))
		toSquirrel := math.Abs(float64(nuts[i][0]-squirrel[0])) + math.Abs(float64(nuts[i][1]-squirrel[1]))
		res += int(2 * toTree)
		if toTree-toSquirrel > maxSelect {
			firstSelect = i
			maxSelect = toTree - toSquirrel
			distanceSelect = toSquirrel
			distanceToTreeSelect = toTree
		}
	}

	if firstSelect >= 0 {
		res = res - int(distanceToTreeSelect) + int(distanceSelect)
	}
	return res
}
