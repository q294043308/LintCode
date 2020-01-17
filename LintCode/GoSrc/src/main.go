package main

import (
	"LogicFun"
)

func main() {
	// head := &Common.TreeNode{Val: 1}
	// head.Left = &Common.TreeNode{Val: 2}
	// head.Right = &Common.TreeNode{Val: 3}
	// head.Left.Left = &Common.TreeNode{Val: 3}
	// head.Left.Right = &Common.TreeNode{Val: 4}
	// head.Right.Left = &Common.TreeNode{Val: 15}
	// head.Right.Right = &Common.TreeNode{Val: 7}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	println(LogicFun.SingleNumberV2Opt([]int{2, 2, 3, 2}))
}
