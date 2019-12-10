package main

import (
	"Common"
	"LogicFun"
)

func main() {
	head := &Common.TreeNode{Val: 2}
	head.Left = &Common.TreeNode{Val: 1}
	head.Right = &Common.TreeNode{Val: 3}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	// head.Next.Next.Next.Next = &Common.ListNode{Val: 5, Next: nil}
	println(LogicFun.IsValidBST(head))
}
