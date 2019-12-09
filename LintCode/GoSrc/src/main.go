package main

import (
	"LogicFun"
)

func main() {
	// head := &Common.ListNode{Val: 1, Next: nil}
	// head.Next = &Common.ListNode{Val: 2, Next: nil}
	// head.Next.Next = &Common.ListNode{Val: 3, Next: nil}
	// head.Next.Next.Next = &Common.ListNode{Val: 4, Next: nil}
	// head.Next.Next.Next.Next = &Common.ListNode{Val: 5, Next: nil}
	println(LogicFun.GenerateTrees(3))
}
