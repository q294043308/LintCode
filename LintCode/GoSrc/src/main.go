package main

import (
	"LogicFun"
)

func main() {
	// head := &Common.TreeNode{Val: 3}
	// head.Left = &Common.TreeNode{Val: 9}
	// head.Right = &Common.TreeNode{Val: 20}
	// head.Right.Left = &Common.TreeNode{Val: 15}
	// head.Right.Right = &Common.TreeNode{Val: 7}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	// head.Next.Next.Next.Next = &Common.ListNode{Val: 5, Next: nil}
	LogicFun.BuildTreeV2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}
