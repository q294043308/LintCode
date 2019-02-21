package main

import "LogicFun"

func main() {
	matrix := make([][]int, 4)
	matrix[0] = []int{1, 1, 1, 1, 1}
	matrix[1] = []int{0, 0, 1, 0, 1}
	matrix[2] = []int{0, 0, 1, 0, 1}
	matrix[3] = []int{0, 0, 1, 0, 0}
	println(LogicFun.RemoveOne(matrix, 1, 2))
}
