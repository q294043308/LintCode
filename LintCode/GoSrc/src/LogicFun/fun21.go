// keep studying(py is a special language,i don't love with it) -- 2020.05.13

package logicfun

import (
	"Common"
	"sort"
	"strconv"
	"strings"
)

// 160. Intersection of Two Linked Lists
func GetIntersectionNode(headA, headB *Common.ListNode) *Common.ListNode {
	dict := make(map[*Common.ListNode]struct{})
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := dict[headA]; ok {
				return headA
			} else {
				dict[headA] = struct{}{}
			}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := dict[headB]; ok {
				return headB
			} else {
				dict[headB] = struct{}{}
			}
			headB = headB.Next
		}
	}
	return nil
}

// 162. Find Peak Element
func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}

	res := 1
	for res < len(nums) {
		rcode := 0
		if res == len(nums)-1 {
			rcode = 1
		} else if nums[res] > nums[res+1] {
			rcode = 2
		} else {
			rcode = 3
		}

		if nums[res] > nums[res-1] && (rcode == 1 || rcode == 2) {
			break
		}
		if rcode == 2 {
			res += 2
		} else {
			res++
		}
	}
	return res
}

// 164. Maximum Gap
func MaximumGap(nums []int) int {
	// 3,6,9,1 -> 3
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

// 165. Compare Version Numbers
func CompareVersion(version1 string, version2 string) int {
	// 1.0.0.0001
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		vv1, _ := strconv.Atoi(v1[i])
		vv2, _ := strconv.Atoi(v2[i])
		if vv1 > vv2 {
			return 1
		}
		if vv1 < vv2 {
			return -1
		}
	}

	v3 := v1
	res := 1
	if i == len(v1) {
		v3 = v2
		res = -1
	}

	for ; i < len(v3); i++ {
		vv3, _ := strconv.Atoi(v3[i])
		if vv3 != 0 {
			return res
		}
	}

	return 0
}

// 167. Two Sum II - Input array is sorted
func TwoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	left := 1
	right := 2

	for true {
		if numbers[left-1]+numbers[right-1] > target {
			return nil
		}
		if numbers[left-1]+numbers[right-1] == target {
			break
		}
		if right < len(numbers) && (numbers[left-1]+numbers[right] <= target) {
			right++
		} else {
			left++
		}
	}

	return []int{left, right}
}

// 166. Fraction to Recurring Decimal
func FractionToDecimal(numerator int, denominator int) string {
	// assert denominator != 0
	res := ""
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		numerator *= -1
		res += "-"
	}
	red := (numerator % denominator) * 10
	res += strconv.Itoa(numerator / denominator)
	if red != 0 {
		res += "."
	}
	numMap := make(map[int]int)
	index := len(res)

	for red != 0 {
		numMap[red] = index
		quo := red / denominator
		index++
		res += strconv.Itoa(quo)
		red %= denominator
		red *= 10

		if v, ok := numMap[red]; ok {
			res = res[:v] + "(" + res[v:]
			res += ")"
			break
		}
	}

	return res
}
