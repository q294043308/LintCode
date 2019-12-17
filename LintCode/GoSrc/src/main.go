package main

import (
	"Common"
	"LogicFun"
	"regexp"
)

func main() {
	head := &Common.TreeNode{Val: 3}
	head.Left = &Common.TreeNode{Val: 9}
	head.Right = &Common.TreeNode{Val: 20}
	head.Right.Left = &Common.TreeNode{Val: 15}
	head.Right.Right = &Common.TreeNode{Val: 7}
	// head.Left.Right.Left = &Common.TreeNode{Val: 31}
	LogicFun.BuildTreeV2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})

	match, _ := regexp.MatchString("d+", "12345678")
	println(match)
	match, _ = regexp.MatchString("d+", "123d")
	println(match)
}
