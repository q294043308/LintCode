package main

import (
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
	a := LogicFun.AddBinary("1", "111")
	println(a)
}
