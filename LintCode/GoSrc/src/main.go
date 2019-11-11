package main

import (
	"Common"
	"LogicFun"
)

type aaa struct {
	num1 int
}

func (self *aaa) Test() {
	println("aaa.test")
}

type cc struct {
	aaa
	int
}

func (self *cc) Test() {
	println("cc.test")
}

func main() {
	head := &Common.ListNode{Val: 1}
	head.Next = &Common.ListNode{Val: 2}
	head.Next.Next = &Common.ListNode{Val: 3}
	println(LogicFun.RotateRight(head, 1))
}
