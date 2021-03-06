// the first letter of function must be bigger, beacause of the special of go-language, I want debug the function in main.go
// so ..  I lose my love .. fighting ~  ^v^

package logicfun

import (
	"Common"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1282. Reverse Vowels of a String
func ReverseVowels(s string) string {
	resBytes := []byte(s)
	var charArr [Common.ENGLISH_CHAR_NUM]bool
	charArr['a'-'a'] = true
	charArr['A'-'A'+Common.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['e'-'a'] = true
	charArr['E'-'A'+Common.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['i'-'a'] = true
	charArr['I'-'A'+Common.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['o'-'a'] = true
	charArr['O'-'A'+Common.SMALL_ENGLISH_CHAR_NUM] = true
	charArr['u'-'a'] = true
	charArr['U'-'A'+Common.SMALL_ENGLISH_CHAR_NUM] = true

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
	var charNumArr [Common.SMALL_ENGLISH_CHAR_NUM]int
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

// 1231. Minimum Moves to Equal Array Elements
func MinMoves(nums []int) int {
	numSum := 0
	minNum := Common.MAXINTNUM
	for i := 0; i < len(nums); i++ {
		numSum = numSum + nums[i]
		minNum = int(math.Min(float64(minNum), float64(nums[i])))
	}

	return numSum - int(minNum)*len(nums)
}

// 1157. Shortest Unsorted Continuous Subarray
func FindUnsortedSubarray(nums []int) int {
	haveFind := false
	minIndex := 0
	minNum := int(Common.MAXINTNUM)
	startIndex := 0
	maxIndex := len(nums) - 1
	maxNum := int(-Common.MININTNUM)
	endIndex := len(nums) - 1

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if nums[i] < minNum {
				minNum = nums[i]
				minIndex = i
				haveFind = true
			}
		}
	}

	if minIndex != 0 {
		for j := minIndex - 1; j >= 0; j-- {
			if nums[minIndex] >= nums[j] {
				startIndex = j + 1
				break
			}
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			if nums[i] > maxNum {
				maxNum = nums[i]
				maxIndex = i
				haveFind = true
			}
		}
	}

	if maxIndex != len(nums)-1 {
		for j := maxIndex + 1; j < len(nums); j++ {
			if nums[maxIndex] <= nums[j] {
				endIndex = j - 1
				break
			}
		}
	}

	if !haveFind {
		return 0
	}
	if endIndex > startIndex {
		return endIndex - startIndex + 1
	}
	return 0
}

// 1237. Number of Boomerangs ---- oh no it's error. i need go home so push it in advance -- i'm coming
func NumberOfBoomerangs(points [][]int) int {
	result := 0

	for i := 0; i < len(points); i++ {
		numMap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			distanceDob := (points[j][0]-points[i][0])*(points[j][0]-points[i][0]) + (points[j][1]-points[i][1])*(points[j][1]-points[i][1])
			numMap[distanceDob]++
		}

		for distance := range numMap {
			result += numMap[distance] * (numMap[distance] - 1)
		}
	}

	return result
}

// 1163. Distribute Candies
func DistributeCandies(candies []int) int {
	sort.Ints(candies)
	numMap := make(map[int][]int)
	temp := candies[0]
	num := 1
	typeNum := 1
	for i := 1; i < len(candies); i++ {
		if temp != candies[i] {
			numMap[num] = append(numMap[num], temp)
			temp = candies[i]
			num = 1
			typeNum++
		} else {
			num++
		}
	}
	numMap[num] = append(numMap[num], temp)

	if len(numMap[1]) >= len(candies)/2 || typeNum >= len(candies)/2 {
		if len(candies)%2 == 0 {
			return len(candies) / 2
		} else {
			return len(candies)/2 + 1
		}
	}
	return typeNum
}

// 1193. Detect Capital
func DetectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	if word[0] <= 'Z' && word[0] >= 'A' {
		var isCapitals bool
		if word[1] <= 'Z' && word[1] >= 'A' {
			isCapitals = true
		} else {
			isCapitals = false
		}
		for i := 2; i < len(word); i++ {
			if isCapitals && word[i] <= 'z' && word[i] >= 'a' {
				return false
			}
			if !isCapitals && word[i] <= 'Z' && word[i] >= 'A' {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] <= 'Z' && word[i] >= 'A' {
				return false
			}
		}
		return true
	}
}

type findShortestSubArraytype struct {
	num        int
	startIndex int
	endIndex   int
}

// 1078. Degree of an Array
func FindShortestSubArray(nums []int) int {
	numIndexMap := make(map[int]*findShortestSubArraytype)

	for i := 0; i < len(nums); i++ {
		if val, ok := numIndexMap[nums[i]]; ok {
			val.num++
			val.endIndex = i
		} else {
			numIndexMap[nums[i]] = &findShortestSubArraytype{
				num:        1,
				startIndex: i,
				endIndex:   i,
			}
		}
	}

	maxNum := int(Common.MININTNUM)
	res := int(Common.MAXINTNUM)
	for _, value := range numIndexMap {
		if maxNum <= value.num {
			maxNum = value.num
			res = int(math.Min(float64(res), float64(value.endIndex-value.startIndex+1)))
		}
	}

	return res
}

// 1319. Contains Duplicate II
func ContainsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	if k == 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if index, ok := numMap[nums[i]]; ok {
			if i-index <= k {
				return true
			}
		}
		numMap[nums[i]] = i
	}
	return false
}

// 1085. Longest Univalue Path ---- LintCode Server is down.. tomorrow is weekend  老子先走为敬，不难为你服务器了
func LongestUnivaluePath(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	longestUnivaluePathSub(root, &res)
	return res
}

func longestUnivaluePathSub(root *Common.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := longestUnivaluePathSub(root.Left, res)
	right := longestUnivaluePathSub(root.Right, res)

	if root.Left != nil && root.Left.Val == root.Val {
		left++
	} else {
		left = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right++
	} else {
		right = 0
	}

	*res = int(math.Max(float64(*res), float64(left+right)))
	return int(math.Max(float64(left), float64(right)))
}

// 993. Array Partition I -- i so sorry that i play game in total morning，But i found the bug of Go IDE in LintCode then i request it;
func ArrayPairSum(nums []int) int {
	res := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		res = res + nums[i]
	}
	return res
}

// 1285. Power of Four
func IsPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num%4 != 0 {
			return false
		}
		num = num / 4
	}
	return true
}

// 1112. Set Mismatch
func FindErrorNums(nums []int) []int {
	res := make([]int, 2)
	sort.Ints(nums)
	lastNum := nums[0]
	isFindDouble := false
	isFindMiss := false

	for i := 1; i < len(nums); i++ {
		if nums[i]-lastNum == 2 {
			res[1] = nums[i] - 1
			isFindMiss = true
		}
		if !isFindDouble && nums[i] == lastNum {
			isFindDouble = true
			res[0] = lastNum
		} else {
			lastNum = nums[i]
		}
	}

	if !isFindMiss {
		res[1] = len(nums)
	}
	if nums[0] != 1 {
		res[1] = 1
	}
	return res
}

// 1253. Convert a Number to Hexadecimal
func ToHex(num int) string {
	var res string
	var isPositive bool
	if num >= 0 {
		isPositive = true
	} else {
		isPositive = false
		num = -num
	}

	for num != 0 {
		Remainder := num % 16
		num = num / 16
		if Remainder >= 10 {
			res = string(Remainder-10+'a') + res
		} else {
			res = string(Remainder-0+'0') + res
		}
	}
	if !isPositive {
		binary := make([]int, 32)
		binary[0] = 1

		for i := len(res) - 1; i >= 0; i-- {
			curChar := res[i]
			var curInt int
			curIndex := 31 - 4*(len(res)-1-i)
			if curChar <= '9' && curChar >= '0' {
				curInt = int(curChar - '0')
			} else {
				curInt = int(curChar-'a') + 10
			}

			for curInt != 0 {
				if curInt%2 == 0 {
					binary[curIndex] = 0
				} else {
					binary[curIndex] = 1
				}
				curIndex--
				curInt = curInt / 2
			}
		}

		for i := 31; i >= 0; i-- {
			if binary[i] == 1 {
				binary[i] = 0
				break
			} else {
				binary[i] = 1
			}
		}
		res = ""
		for i := 1; i < 32; i++ {
			if binary[i] == 1 {
				binary[i] = 0
			} else {
				binary[i] = 1
			}
		}
		for i := 0; i < 32; i = i + 4 {
			val := 8*binary[i] + 4*binary[i+1] + 2*binary[i+2] + binary[i+3]
			if val >= 10 {
				res = res + string(val-10+'a')
			} else {
				res = res + string(val-0+'0')
			}
		}
	}

	if len(res) == 0 {
		res = "0"
	}
	return res
}

// 1199. Perfect Number
func CheckPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	res := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			if num/i == i {
				res += i
			} else {
				res += i + num/i
			}
		}

		if res > num {
			return false
		}
		if res == num {
			return true
		}
	}
	if res == num {
		return true
	}
	return false
}

// 1104. Judge Route Circle
func JudgeCircle(moves string) bool {
	point := &Common.Point{
		X: 0,
		Y: 0,
	}

	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'R':
			point.X++
		case 'L':
			point.X--
		case 'U':
			point.Y++
		case 'D':
			point.Y--
		}
	}

	if point.X == 0 && point.Y == 0 {
		return true
	}
	return false
}

// 1320. Contains Duplicate
func ContainsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			return true
		} else {
			numMap[nums[i]] = true
		}
	}
	return false
}

// 1294. Power of Three
func IsPowerOfThree(n int) bool {
	/* this solution is from math, but the result of Log10 func isn't precisely, so it's error
		if n > 0 && (float64(int(math.Log10(float64(n))/math.Log10(3)))-(math.Log10(float64(n))/math.Log10(3)) == 0) {
			return true
		}
	    return false
	*/

	// here is a foolish func, because of the limit of 'int' type, so n is litter by 3 ^ 19 -->>
	maxIntForPowerOfThree := int(math.Pow(3, 19))
	if maxIntForPowerOfThree%n == 0 {
		return true
	}
	return false
}

// 1218. Number Complement
func FindComplement(num int) int {
	var bitStr string
	var res int
	curBase := 1

	for num > 0 {
		if num%2 == 0 {
			bitStr = "1" + bitStr
		} else {
			bitStr = "0" + bitStr
		}
		num /= 2
	}

	for i := len(bitStr) - 1; i >= 0; i-- {
		if bitStr[i] == '1' {
			res += curBase
		}
		curBase *= 2
	}

	return res
}

// 977. Base 7
func ConvertToBase7(num int) string {
	var res string
	isPossive := true
	if num < 0 {
		isPossive = false
		num = -num
	}

	for num > 0 {
		res = string('0'+num%7) + res
		num /= 7
	}

	if len(res) == 0 {
		res = "0"
	}
	if !isPossive {
		res = "-" + res
	}

	return res
}

// 1243. Number of Segments in a String
func CountSegments(s string) int {
	res := 0
	lastSpace := true

	for _, val := range s {
		if val != ' ' {
			if lastSpace {
				res++
			}
			lastSpace = false
		} else {
			lastSpace = true
		}
	}

	return res
}

// 916. Palindrome Permutation
func CanPermutePalindrome(s string) bool {
	var charNum [Common.CHARNUM]int
	count := 0

	for _, val := range s {
		if charNum[val] == 1 {
			charNum[val] = 0
		} else {
			charNum[val] = 1
		}
	}

	for _, val := range charNum {
		if val == 1 {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}

// 1178. Student Attendance Record I
func CheckRecord(s string) bool {
	Acount := 0
	Lcount := 0

	for _, val := range s {
		if val == 'L' {
			if Lcount >= 2 {
				return false
			}
			Lcount++
		} else {
			Lcount = 0
			if val == 'A' {
				if Acount >= 1 {
					return false
				}
				Acount = 1
			}
		}
	}
	return true
}

// 1099. Non-decreasing Array
func CheckPossibility(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	isModify := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if isModify {
				return false
			} else {
				if i+2 < len(nums) && nums[i+2] < nums[i] {
					if i != 0 {
						return false
					}
				}
				isModify = true
			}
		}
	}
	return true
}

// 1300. Nim Game -- i'm a man of genius
func CanWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}

	return true
}

// 644. Strobogrammatic Number
func IsStrobogrammatic(num string) bool {
	middle := 0
	if len(num)%2 == 0 {
		middle = len(num) / 2
	} else {
		middle = (len(num) + 1) / 2
	}
	for i := 0; i < middle; i++ {
		if (num[i] == '1' && num[len(num)-1-i] == '1') ||
			(num[i] == '8' && num[len(num)-1-i] == '8') ||
			(num[i] == '0' && num[len(num)-1-i] == '0') {
			continue
		} else if num[i] == '6' && num[len(num)-1-i] == '9' {
			continue
		} else if num[i] == '9' && num[len(num)-1-i] == '6' {
			continue
		}
		return false
	}

	return true
}

// 647. Find All Anagrams in a String
func FindAnagrams(s string, p string) []int {
	res := make([]int, 0, 10)
	i := 0
	targetMap := make(map[int]int)
	curMap := make(map[int]int)

	for _, val := range p {
		targetMap[int(val)]++
	}
	for ; i < len(s) && i < len(p); i++ {
		curMap[int(s[i])]++
	}

	for ; i <= len(s); i++ {
		isEqual := true
		for key, val := range targetMap {
			if _, ok := curMap[key]; !ok {
				isEqual = false
				break
			}
			if val != curMap[key] {
				isEqual = false
				break
			}
		}
		if isEqual {
			res = append(res, i-len(p))
		}

		if i < len(s) {
			curMap[int(s[i-len(p)])]--
			curMap[int(s[i])]++
		}
	}

	return res
}

var TwoSumMap = make(map[int]int)

// 607. Two Sum III - Data structure design
func TwoSumAdd(number int) {
	TwoSumMap[number]++
}

func TwoSumFind(value int) bool {
	for num, _ := range TwoSumMap {
		if value-num == num {
			if TwoSumMap[num] >= 2 {
				return true
			}
		} else if TwoSumMap[value-num] >= 1 {
			return true
		}

	}
	return false
}

// 637. Valid Word Abbreviation
func ValidWordAbbreviation(word string, abbr string) bool {
	i := 0
	j := 0

	for i < len(word) {
		curIndex := 0
		for ; j < len(abbr); j++ {
			if abbr[j] <= '9' && abbr[j] >= '0' {
				curIndex = curIndex*10 + int(abbr[j]-'0')
				if curIndex == 0 {
					break
				}
			} else {
				break
			}
		}
		if curIndex == 0 {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		} else {
			i = i + curIndex
		}
	}

	if i == len(word) && j == len(abbr) {
		return true
	}
	return false
}

// 626. Rectangle Overlap
func DoOverlap(l1 *Common.Point, r1 *Common.Point, l2 *Common.Point, r2 *Common.Point) bool {
	if math.Max(float64(l1.X), float64(l2.X)) <= math.Min(float64(r1.X), float64(r2.X)) && math.Min(float64(l1.Y), float64(l2.Y)) >= math.Max(float64(r1.Y), float64(r2.Y)) {
		return true
	}
	return false
}

func longestConsecutiveSub(root *Common.TreeNode, res *int, curLen int) {
	if root.Left != nil {
		if root.Left.Val == root.Val+1 {
			longestConsecutiveSub(root.Left, res, curLen+1)
		} else {
			longestConsecutiveSub(root.Left, res, 1)
		}
	}
	if root.Right != nil {
		if root.Right.Val == root.Val+1 {
			longestConsecutiveSub(root.Right, res, curLen+1)
		} else {
			longestConsecutiveSub(root.Right, res, 1)
		}
	}

	if curLen > *res {
		*res = curLen
	}
}

// 595. Binary Tree Longest Consecutive Sequence
func LongestConsecutive(root *Common.TreeNode) int {
	res := 0

	if root == nil {
		return res
	}

	longestConsecutiveSub(root, &res, 1)
	return res
}

// 643. Longest Absolute File Path
func LengthLongestPath(input string) int {
	var maxLen int
	var pathArr []string
	startIndex := 0
	curPath := ""
	curIndex := 0

	if len(input) > 0 {
		baseIndex := strings.Index(input[startIndex:], "\\n")
		if baseIndex == -1 {
			return 0
		}
		pathArr = append(pathArr, input[0:baseIndex])
		startIndex = baseIndex + 2
	} else {
		return 0
	}

	for startIndex < len(input) {
		endIndex := strings.Index(input[startIndex:], "\\n") + startIndex
		if endIndex < startIndex {
			endIndex = len(input)
		}

		tCount := 0
		for i := startIndex; i+1 < len(input); i += 2 {
			if input[i] == '\\' && input[i+1] == 't' {
				tCount++
			} else {
				break
			}
		}
		if tCount == curIndex+1 {
			pathArr = append(pathArr, input[startIndex+tCount*2:endIndex])
			curIndex++
		} else {
			if strings.Count(pathArr[len(pathArr)-1], ".") > 0 {
				curPath = ""
				for i := 0; i < len(pathArr); i++ {
					curPath += pathArr[i]
					curPath += "/"
				}
				if len(curPath)-1 > maxLen {
					maxLen = len(curPath) - 1
				}
			}
			pathArr = append(pathArr[0:tCount], input[startIndex+tCount*2:endIndex])
		}
		startIndex = endIndex + 2
	}

	if strings.Count(pathArr[len(pathArr)-1], ".") > 0 {
		curPath = ""
		for i := 0; i < len(pathArr); i++ {
			curPath += pathArr[i]
			curPath += "/"
		}
		if len(curPath)-1 > maxLen {
			maxLen = len(curPath) - 1
		}
	}
	return maxLen
}

// 575. Decode String
func ExpressionExpand(s string) string {
	start := 0
	res := ""

	for strings.Index(s[start:], "[") >= 0 {
		leftCount := 1
		curStart := strings.Index(s[start:], "[") + start
		curNum := ""
		j := curStart - 1
		for ; j >= 0; j-- {
			if s[j] > '9' || s[j] < '0' {
				break
			}
			curNum = string(s[j]) + curNum
		}

		res += s[start : j+1]
		curEnd := curStart + 1
		for ; curEnd < len(s); curEnd++ {
			if s[curEnd] == '[' {
				leftCount++
			}
			if s[curEnd] == ']' {
				leftCount--
				if leftCount == 0 {
					break
				}
			}
		}
		tmpStr := ExpressionExpand(s[curStart+1 : curEnd])
		num, _ := strconv.Atoi(curNum)
		for j = 0; j < num; j++ {
			res += tmpStr
		}

		start = curEnd + 1
	}

	res += s[start:]
	return res
}

func FindNumberOfLISError(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	resLen := 0
	resNum := 0
	curLen := 1

	for i := 1; i <= len(nums); i++ {
		if i != len(nums) && nums[i] > nums[i-1] {
			curLen++
			continue
		} else if curLen == resLen {
			resNum++
			continue
		} else if curLen > resLen {
			resLen = curLen
			resNum = 1
			continue
		}
	}
	return resNum
}

// 1093. Number of Longest Increasing Subsequence (Dynamic Program)
func FindNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	type LISNode struct {
		len int
		num int
	}

	// 1, 3, 5, 4, 7
	resLen := -1
	onceNum := 0
	maxNumArr := make([]LISNode, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				if maxNumArr[i].len < maxNumArr[j].len+1 {
					maxNumArr[i].len = maxNumArr[j].len + 1
					maxNumArr[i].num = maxNumArr[j].num
					resLen = i
				} else if maxNumArr[i].len == maxNumArr[j].len+1 {
					maxNumArr[i].num += maxNumArr[j].num
				}
			}
		}
		if maxNumArr[i].len == 0 {
			maxNumArr[i].len = 1
			maxNumArr[i].num = 1
			onceNum++
		}

	}

	if resLen == -1 {
		return onceNum
	}
	return maxNumArr[resLen].num
}

// 5. Kth Largest Element
func KthLargestElement(n int, nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-n]
}
