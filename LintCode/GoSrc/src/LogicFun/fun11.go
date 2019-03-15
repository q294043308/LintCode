// There have a lot of bugs in LintCode; i feeled like eating shit！
// escape LintCode, Go LeetCode -- 2019.3.7
package LogicFun

import (
	"Common"
	"math"
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

// 4. Longest Palindromic Substring (it's haven a Algorithms of malacher, Limited to O(n) time.)
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
