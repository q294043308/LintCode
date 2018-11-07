package main

import (
	"CommonFun"
	"LogicFun"
	"fmt"
)

func main() {
	s := new(CommonFun.TreeNode)
	s.Val = 1
	s.Right = new(CommonFun.TreeNode)
	s.Right.Val = 3
	s.Right.Left = new(CommonFun.TreeNode)
	s.Right.Left.Val = 2
	s.Right.Right = new(CommonFun.TreeNode)
	s.Right.Right.Val = 4
	s.Right.Right.Right = new(CommonFun.TreeNode)
	s.Right.Right.Right.Val = 5
	fmt.Println(LogicFun.LongestConsecutive(s))
}
