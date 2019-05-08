package main

import (
	"LogicFun"
)

func main() {
	res := LogicFun.IsMatch("mississippi", "mis*is*p*ip*.") // true
	println(res)
	res = LogicFun.IsMatch("mississippi", "mis*is*ip*.") // true
	println(res)
	res = LogicFun.IsMatch("mississippi", "mis*is*p*.") //false
	println(res)
	res = LogicFun.IsMatch("a", "ab*a") // false
	println(res)
	res = LogicFun.IsMatch("aab", "c*a*b") // true
	println(res)
	res = LogicFun.IsMatch("abcd", "d*") //false
	println(res)
	res = LogicFun.IsMatch("aaa", "a*aaa") // true
	println(res)
	res = LogicFun.IsMatch("aaa", "a*aaaa") //false
	println(res)
	res = LogicFun.IsMatch("aaaaaaaaab", "a*a*a*b") //true
	println(res)
	res = LogicFun.IsMatch("aaa", "ab*a*c*a") //true
	println(res)
	res = LogicFun.IsMatch("a", ".*..a*") //false
	println(res)
	res = LogicFun.IsMatch("ab", ".*..c*") //true
	println(res)
	res = LogicFun.IsMatch("cbaacacaaccbaabcb", "c*b*b*.*ac*.*bc*a*") //true
	println(res)
}
