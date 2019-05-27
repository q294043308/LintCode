// There have a lot of bugs in LintCode; i feeled like eating shit！
// escape LintCode, Go LeetCode -- 2019.3.7
package LogicFun

import (
	"Common"
	"errors"
	"math"
	"sort"
	"strings"
)

// 1. Two Sum
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	res := make([]int, 2)

	for in, num := range nums {
		if _, ok := numMap[target-num]; ok {
			res[0] = numMap[target-num]
			res[1] = in
			return res
		} else {
			numMap[num] = in
		}
	}

	return res
}

// 2. Add Two Numbers
func AddTwoNumbers(l1 *Common.ListNode, l2 *Common.ListNode) *Common.ListNode {
	var isAdd byte = 0
	var l1Last *Common.ListNode
	res := l1

	for l1 != nil && l2 != nil {
		l1.Val += l2.Val
		if isAdd == 1 {
			l1.Val++
		}

		if l1.Val >= 10 {
			l1.Val -= 10
			isAdd = 1
		} else {
			isAdd = 0
		}
		l1Last = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		l1Last.Next = l2
		l1 = l1Last.Next
	}

	for l1 != nil {
		if isAdd == 1 {
			l1.Val++
			if l1.Val >= 10 {
				l1.Val -= 10
				l1Last = l1
				l1 = l1.Next
				continue
			}
		}
		return res
	}

	if isAdd == 1 {
		l1Last.Next = &Common.ListNode{
			Val: 1,
		}
	}
	return res
}

// 3. Longest Substring Without Repeating Characters
func LengthOfLongestSubstring(s string) int {
	charMap := make([]int, Common.CHARNUM, Common.CHARNUM)
	res := 0
	lastIndex := 0

	for i := 0; i < len(s); i++ {
		if charMap[s[i]] > lastIndex {
			lastIndex = charMap[s[i]]
		}

		if res < i-lastIndex+1 {
			res = i - lastIndex + 1
		}

		charMap[s[i]] = i + 1
	}

	return res
}

// 4. Median of Two Sorted Arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	if nums1Len > nums2Len {
		// exchange
		temp := nums1
		nums1 = nums2
		nums2 = temp

		nums1Len = len(nums1)
		nums2Len = len(nums2)
	}

	iMin := 0
	iMax := nums1Len
	halfLen := (nums1Len + nums2Len + 1) / 2

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i

		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}
			if (nums1Len+nums2Len)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == nums1Len {
				minRight = nums2[j]
			} else if j == nums2Len {
				minRight = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					minRight = nums1[i]
				} else {
					minRight = nums2[j]
				}
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

func FindMedianSortedArraysSub(nums []int) (float64, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("empty")
	}

	index := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[index]+nums[index-1]) / 2, index - 1, nil
	} else {
		return float64(nums[index]), index, nil
	}
}

// 5. Longest Palindromic Substring (it's haven a Algorithms of malacher, Limited to O(n) time.)
func LongestPalindrome(s string) string {
	start, end, le := 0, 0, len(s)
	if le == 0 {
		return ""
	}
	for i := 0; i < le; i++ {
		if end-start > (le-i)*2 {
			return s[start : end+1]
		}

		len1 := expand(s, i, i+1)
		len2 := expand(s, i, i)
		var len int
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + (len)/2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

// malacher
func LongestPalindromeMaLaCher(s string) string {
	len := len(s)
	if len == 0 {
		return ""
	}

	news := make([]byte, len*2+3)
	p := make([]int, len*2+3)
	news[0] = '~'
	news[1] = '!'
	j := 2

	for i := 0; i < len; i++ {
		news[j] = s[i]
		news[j+1] = '!'
		j += 2
	}
	news[j] = '@'

	mx, id := 0, 0
	maxLen, maxIndex := 0, 0
	for i := 1; i < len*2+2; i++ {
		if mx > i {
			if p[2*id-i] < mx-i {
				p[i] = p[2*id-i]
			} else {
				p[i] = mx - i
			}
		} else {
			p[i] = 1
		}
		for news[i+p[i]] == news[i-p[i]] {
			p[i]++
		}
		if p[i]+i > mx {
			mx = p[i] + i
			id = i
		}

		if p[i] >= maxLen {
			maxLen = p[i]
			maxIndex = i
		}
	}

	maxLen -= 1
	if maxIndex%2 == 0 {
		maxIndex = (maxIndex - 2) / 2
		return s[maxIndex-maxLen/2 : maxIndex-maxLen/2+maxLen]
	} else {
		maxIndex = (maxIndex-2)/2 + 1
		return s[maxIndex-maxLen/2 : maxIndex-maxLen/2+maxLen]
	}
}

// 7. Reverse Integer
func Reverse(x int) int {
	var res int64
	for x != 0 {
		res = res*10 + int64(x%10)
		x /= 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return int(res)
}

// 10. Regular Expression Matching
func IsMatch(s string, p string) bool {
	if len(s) == 0 {
		if s == "" && p == "" {
			return true
		}

		num := 0
		for _, val := range p {
			if val == '*' {
				if num <= 0 {
					return false
				}
				num--
			} else {
				num++
			}
		}
		if num == 0 {
			return true
		}

		return false
	}
	if len(p) == 0 || p[0] == '*' {
		return false
	}

	matchMarix := make([][]bool, len(p))
	for i := 0; i < len(p); i++ {
		matchMarix[i] = make([]bool, len(s))
	}

	lastJStart := 0
	jStart := 0
	trueNumStart := 1

	for i := 0; i < len(p); i++ {
		isUpdate := false
		j := 0

		if i > 0 {
			if p[i] != '*' {
				if p[i-1] == '*' {
					if trueNumStart > jStart {
						j = trueNumStart
					} else {
						j = jStart
					}
				} else {
					j = jStart + 1
				}
				trueNumStart++
			} else {
				j = jStart
				trueNumStart--
			}
		}

		if i > 1 && p[i] == '*' {
			for z := lastJStart; z < len(s); z++ {
				if matchMarix[i-2][z] {
					matchMarix[i][z] = matchMarix[i-2][z]
					if !isUpdate {
						isUpdate = true
						lastJStart = jStart
						jStart = z
					}
				}
			}
		}

		for ; j < len(s); j++ {
			if s[j] == p[i] || p[i] == '.' {
				if i == 0 || j == 0 || matchMarix[i-1][j-1] {
					matchMarix[i][j] = true

					if !isUpdate {
						isUpdate = true
						lastJStart = jStart
						jStart = j
						if trueNumStart < jStart {
							trueNumStart = jStart
						}
					}
				}

				if i == 0 {
					break
				}
			} else if p[i] == '*' {
				if i == 0 {
					return false
				}

				if (s[j] == p[i-1] || p[i-1] == '.') && (matchMarix[i-1][j] || j > 0 && matchMarix[i][j-1]) {
					matchMarix[i][j] = true

					if !isUpdate {
						isUpdate = true
						lastJStart = jStart
						jStart = j
					}
				}
			} else if i > 0 && !matchMarix[i-1][j] {
				if !isUpdate {
					isUpdate = true
					lastJStart = jStart
					jStart = j
				}
			} else if i == 0 {
				break
			}
		}
		if !isUpdate {
			isUpdate = true
			lastJStart = jStart
		}
	}

	return matchMarix[len(p)-1][len(s)-1]
}

// 11. Container With Most Water
func MaxArea(height []int) int {
	res := 0
	compIndexMap := make(map[int]struct{})
	lastIndex := -1

	for index, h := range height {
		if index < len(height)-1 && height[index+1] >= height[index] {
			compIndexMap[index] = struct{}{}
			lastIndex = index
			continue
		}
		for preIndex, _ := range compIndexMap {
			preHei := height[preIndex]
			hei := h
			if preHei < h {
				hei = preHei
			}
			curVal := (index - preIndex) * hei
			if curVal > res {
				res = curVal
			}
		}

		if index > lastIndex {
			compIndexMap[index] = struct{}{}
			lastIndex = index
		}
	}

	return res
}

// 11+. Container With Most Water
func MaxAreaOpt(height []int) int {
	res := 0
	leftIndex := 0
	rightIndex := len(height) - 1

	for rightIndex > leftIndex {
		het := height[leftIndex]
		if height[leftIndex] > height[rightIndex] {
			het = height[rightIndex]
		}

		curVal := (rightIndex - leftIndex) * het
		if res < curVal {
			res = curVal
		}

		if height[rightIndex] > height[leftIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return res
}

// 12. Integer to Roman
func IntToRoman(num int) string {
	res := ""
	romanByte := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	i := 6
	dev := 1000

	for i >= 0 {
		mod := num / dev
		if mod != 0 {
			if mod == 9 {
				res = res + string(romanByte[i]) + string(romanByte[i+2])
			} else if mod == 4 {
				res = res + string(romanByte[i]) + string(romanByte[i+1])
			} else {
				newMod := mod
				if newMod >= 5 {
					newMod -= 5
					res = res + string(romanByte[i+1])
				}
				for num := 0; num < newMod; num++ {
					res = res + string(romanByte[i])
				}

			}
		}

		num -= mod * dev
		dev /= 10
		i -= 2
	}

	return res
}

// 13. Roman to Integer
func RomanToInt(s string) int {
	res := 0
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	lastNum := 10000

	for _, char := range s {
		curVal := romanMap[char]

		if curVal > lastNum {
			res += curVal - 2*lastNum
		} else {
			res += curVal
			lastNum = curVal
		}
	}

	return res
}

// 13+. Roman to Integer
func RomanToIntOpt(s string) int {
	romanArr := make([]int16, 24)
	romanArr['I'-'A'] = 1
	romanArr['V'-'A'] = 5
	romanArr['X'-'A'] = 10
	romanArr['L'-'A'] = 50
	romanArr['C'-'A'] = 100
	romanArr['D'-'A'] = 500
	romanArr['M'-'A'] = 1000
	var lastNum int16 = 10000
	var res int16

	for _, char := range s {
		curVal := romanArr[char-'A']

		if curVal > lastNum {
			res += curVal - 2*lastNum
		} else {
			res += curVal
			lastNum = curVal
		}
	}

	return int(res)
}

// 14. Longest Common Prefix
func LongestCommonPrefix(strs []string) string {
	shortestStr := ""
	shortestIndex := -1
	for index, str := range strs {
		if len(shortestStr) == 0 {
			shortestStr = str
			shortestIndex = index
		} else if len(shortestStr) > len(str) {
			shortestStr = str
			shortestIndex = index
		}
	}

	for lenRange := len(shortestStr); lenRange > 0; lenRange-- {
		strRange := shortestStr[0:lenRange]
		ok := true

		for index, str := range strs {
			if shortestIndex == index {
				continue
			}

			if strings.HasPrefix(str, strRange) {
				continue
			}

			ok = false
			break
		}

		if ok {
			return strRange
		}
	}

	return ""
}

// 15. 3Sum -> this question have a solution of less time, this solution will using at 16's question
func ThreeSum(nums []int) [][]int {
	type repeatStr struct {
		repeatMap map[int]struct{}
		cnt       int
	}

	res := make([][]int, 0)
	numMap := make(map[int]*repeatStr, 0)
	for _, num := range nums {
		if val, ok := numMap[num]; ok {
			val.cnt++
		} else {
			numMap[num] = new(repeatStr)
			numMap[num].repeatMap = make(map[int]struct{})
		}
	}

	for i, iV := range nums {
		for j := i + 1; j < len(nums); j++ {
			jV := nums[j]
			needNum := 0 - iV - jV
			if needNum == iV || needNum == jV {
				if needNum == iV && needNum == jV && numMap[needNum].cnt < 2 {
					continue
				} else if numMap[needNum].cnt < 1 {
					continue
				}
			}

			repeatVal, ok := numMap[needNum]
			if ok {
				if _, repeat := repeatVal.repeatMap[iV]; !repeat {
					res = append(res, []int{iV, jV, needNum})
					repeatVal.repeatMap[iV] = struct{}{}
					repeatVal.repeatMap[jV] = struct{}{}
					numMap[iV].repeatMap[jV] = struct{}{}
					numMap[iV].repeatMap[needNum] = struct{}{}
					numMap[jV].repeatMap[iV] = struct{}{}
					numMap[jV].repeatMap[needNum] = struct{}{}
				}
			}
		}
	}

	return res
}

// 16. 3Sum Closest
func ThreeSumClosest(nums []int, target int) int {
	res := 0
	resDis := Common.MAXINTNUM
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1

		for start != end {
			sum := nums[i] + nums[start] + nums[end]
			dis := 0
			if sum < target {
				dis = target - sum
				start++
			} else {
				dis = sum - target
				end--
			}
			if dis == 0 {
				return target
			}
			if dis < resDis {
				resDis = dis
				res = sum
			}
		}
	}
	return res
}

// 17. Letter Combinations of a Phone Number
func LetterCombinations(digits string) []string {
	letterArr := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	size := 1
	for _, digit := range digits {
		size = size * len(letterArr[digit-'2'])
	}

	if size < 3 {
		size = 0
	}
	res := make([]string, size)
	for _, digit := range digits {
		letter := letterArr[digit-'2']
		size = size / len(letter)
		for i := 0; i < len(res); i++ {
			res[i] = res[i] + string(letter[(i/size)%len(letter)])
		}
	}

	return res
}

/*
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.
*/
// 18. 4Sum
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-3; {
		cur := nums[i]
		for j := i + 1; j < len(nums)-2; {
			cur += nums[j]
			start := j + 1
			end := len(nums) - 1

			for start < end {
				if cur+nums[start]+nums[end] == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
				}

				if cur+nums[start]+nums[end] < target {
					newstart := start + 1
					for newstart < end && nums[newstart] == nums[start] {
						newstart++
					}
					start = newstart
				} else {
					newend := end - 1
					for start < newend && nums[newend] == nums[end] {
						newend--
					}
					end = newend
				}
			}

			cur -= nums[j]
			newj := j + 1
			for newj < len(nums)-2 && nums[newj] == nums[j] {
				newj++
			}
			j = newj
		}

		newi := i + 1
		for newi < len(nums)-3 && nums[newi] == nums[i] {
			newi++
		}
		i = newi
	}

	return res
}
