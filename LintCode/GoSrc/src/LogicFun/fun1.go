// the first letter of function must be bigger, beacause of the special of go-language, I want debug the function in main.go
// so ..  I lose my love .. fighting ~  ^v^

package LogicFun

import (
	"CommonFun"
	"math"
	"strconv"
	"strings"
)

// 1282. Reverse Vowels of a String
func ReverseVowels(s string) string {
	resBytes := []byte(s)
	var charArr [CommonFun.ENGLISH_CHAR_NUM]bool
	charArr['a'-'a'] = true
	charArr['A'-'A'+CommonFun.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['e'-'a'] = true
	charArr['E'-'A'+CommonFun.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['i'-'a'] = true
	charArr['I'-'A'+CommonFun.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['o'-'a'] = true
	charArr['O'-'A'+CommonFun.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['u'-'a'] = true
	charArr['U'-'A'+CommonFun.SMALL_ENGLISH_CHAR_NUM] = true

	startDex := 0
	front := false
	behind := false

	for endDex := len(s) - 1; startDex < endDex; {
		if !front &&
			((s[startDex] <= 'z' && s[startDex] >= 'a' && charArr[s[startDex]-'a']) ||
				(s[startDex] <= 'Z' && s[startDex] >= 'A' && charArr[s[startDex]-'A'])) {
			front = true
		}
		if !behind &&
			((s[endDex] <= 'z' && s[endDex] >= 'a' && charArr[s[endDex]-'a']) ||
				(s[endDex] <= 'Z' && s[endDex] >= 'A' && charArr[s[endDex]-'A'])) {
			behind = true
		}

		if front && behind {
			tmp := resBytes[startDex]
			resBytes[startDex] = resBytes[endDex]
			resBytes[endDex] = tmp
			startDex++
			endDex--
			front = false
			behind = false
		} else {
			if !front {
				startDex++
			}
			if !behind {
				endDex--
			}
		}
	}

	return string(resBytes)
}

// 1385. Lucky Number Eight
func LuckyNumber(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if strings.Contains(strconv.Itoa(i), "8") {
			res++
		}
	}

	return res
}

// 914. Flip Game
func GeneratePossibleNextMoves(s string) []string {
	res := make([]string, 0, 0)
	strByte := []byte(s)
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '+' && s[i-1] == '+' {
			strByte[i-1] = '-'
			strByte[i] = '-'
			res = append(res, string(strByte))
			strByte[i-1] = '+'
			strByte[i] = '+'
		}
	}

	return res
}

// 1314. Power of Two
func IsPowerOfTwo(n int) bool {
	for n >= 1 {
		if n == 2 || n == 1 {
			return true
		}
		if n%2 == 0 {
			n /= 2
		} else {
			return false
		}
	}
	return false
}

// 1228. Poor Pigs
func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	time := minutesToTest/minutesToDie + 1
	res := 0
	for math.Pow(float64(time), float64(res)) < float64(buckets) {
		res++
	}

	return res
}

// 1086. Repeated String Match
func RepeatedStringMatch(A string, B string) int {
	firstIndex := strings.Index(B, A)
	if firstIndex > len(A) {
		return -1
	}

	for i := firstIndex - 1; i >= 0; i-- {
		if B[i] != A[len(A)-firstIndex+i] {
			return -1
		}
	}

	indexA := 0
	for i := firstIndex + len(A); i < len(B); i++ {
		if B[i] != A[indexA] {
			return -1
		}
		indexA = indexA + 1
		if indexA == len(A) {
			indexA = 0
		}
	}

	if firstIndex == -1 {
		if strings.Contains(A, B) {
			return 1
		}
		if strings.Contains(A+A, B) {
			return 2
		}
		return -1
	} else if firstIndex == 0 {
		if len(B)%len(A) == 0 {
			return len(B) / len(A)
		} else {
			return len(B)/len(A) + 1
		}
	} else {
		return 2 + (len(B)-firstIndex)/len(A)
	}
}

// 1270. Ransom Note
func CanConstruct(ransomNote string, magazine string) bool {
	var charNumArr [CommonFun.SMALL_ENGLISH_CHAR_NUM]int
	for i := 0; i < len(magazine); i++ {
		charNumArr[magazine[i]-'a'] = charNumArr[magazine[i]-'a'] + 1
	}
	for i := 0; i < len(ransomNote); i++ {
		if charNumArr[ransomNote[i]-'a'] > 0 {
			charNumArr[ransomNote[i]-'a'] = charNumArr[ransomNote[i]-'a'] - 1
		} else {
			return false
		}
	}
	return true
}

// 1068. Find Pivot Index
func PivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	leftNum := 0
	rightNum := 0
	for i := 1; i < len(nums); i++ {
		rightNum = rightNum + nums[i]
	}

	if leftNum == rightNum {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		leftNum = leftNum + nums[i-1]
		rightNum = rightNum - nums[i]
		if leftNum == rightNum {
			return i + 1
		}
	}
	return -1
}

// 1056. Find Smallest Letter Greater Than Target
func NextGreatestLetter(letters string, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}
	return letters[0]
}

// 973. 1-bit and 2-bit Characters
func IsOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i = i + 2
		} else {
			i = i + 1
		}
	}
	return i == len(bits)-1
}
