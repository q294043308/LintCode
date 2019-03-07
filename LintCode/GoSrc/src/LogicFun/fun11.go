// There have a lot of bugs in LintCode; i feeled like eating shit！
// escape LintCode, Go LeetCode -- 2019.3.7
package LogicFun

import "Common"

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
