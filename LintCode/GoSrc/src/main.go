package main

import (
	"LogicFun"
)

func main() {
	var s int = -2147483648 / -1
	println(s)
	println(LogicFun.Divide(-2147483648, -1))
}
