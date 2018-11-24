package main

import (
	"LogicFun"
	"fmt"
)

func main() {
	var banned []string
	banned = append(banned, "bob")
	banned = append(banned, "hit")
	fmt.Println(LogicFun.MostCommonWord("Bob. hIt, baLl", banned))
}
