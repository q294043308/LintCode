package LogicFun

import (
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
