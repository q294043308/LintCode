package main

import (
	"Common"
	"LogicFun"
)

func main() {
	head := &Common.ListNode{Val: 3}
	head.Next = &Common.ListNode{Val: 2}
	head.Next.Next = &Common.ListNode{Val: 0}
	head.Next.Next.Next = &Common.ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next
	// head.Left = &Common.TreeNode{Val: 2}
	// head.Right = &Common.TreeNode{Val: 3}
	// head.Left.Left = &Common.TreeNode{Val: 3}
	// head.Left.Right = &Common.TreeNode{Val: 4}
	// head.Right.Left = &Common.TreeNode{Val: 15}
	// head.Right.Right = &Common.TreeNode{Val: 7}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	println(LogicFun.HasCycleOpt(head))
}
