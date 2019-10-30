package main

import "LogicFun"

type aaa struct {
	num int
}

func main() {
	// aaa := make([]int, 0, 10)
	// aaa = append(aaa, 10)
	// println(aaa)
	LogicFun.SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
