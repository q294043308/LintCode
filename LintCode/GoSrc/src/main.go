package main

import (
	"LogicFun"
)

func main() {
	// head := &Common.TreeNode{Val: -3}
	// head.Left = &Common.TreeNode{Val: 2}
	// head.Right = &Common.TreeNode{Val: 3}
	// head.Left.Left = &Common.TreeNode{Val: 3}
	// head.Left.Right = &Common.TreeNode{Val: 4}
	// head.Right.Left = &Common.TreeNode{Val: 15}
	// head.Right.Right = &Common.TreeNode{Val: 7}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	LogicFun.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
}
