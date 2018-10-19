// the first letter of function must be bigger, beacause of the special of go-language, I want debug the function in main.go
// so ..  I lose my love .. fighting ~  ^v^

package LogicFun

import (
	"CommonFun"
	"math"
	"sort"
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

// 1231. Minimum Moves to Equal Array Elements
func MinMoves(nums []int) int {
	numSum := 0
	minNum := CommonFun.MAXINTNUM
	for i := 0; i < len(nums); i++ {
		numSum = numSum + nums[i]
		minNum = math.Min(float64(minNum), float64(nums[i]))
	}

	return numSum - int(minNum)*len(nums)
}

// 1157. Shortest Unsorted Continuous Subarray
func FindUnsortedSubarray(nums []int) int {
	haveFind := false
	minIndex := 0
	minNum := int(CommonFun.MAXINTNUM)
	startIndex := 0
	maxIndex := len(nums) - 1
	maxNum := int(-CommonFun.MININTNUM)
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

	maxNum := int(CommonFun.MININTNUM)
	res := int(CommonFun.MAXINTNUM)
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
func LongestUnivaluePath(root *CommonFun.TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	longestUnivaluePathSub(root, &res)
	return res
}

func longestUnivaluePathSub(root *CommonFun.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := longestUnivaluePathSub(root.Left, res)
	right := longestUnivaluePathSub(root.Right, res)

	if root.Right.Val == root.Val {
		left++
	} else {
		left = 0
	}
	if root.Right.Val == root.Val {
		right++
	} else {
		right = 0
	}

	*res = int(math.Max(float64(*res), float64(left+right)))
	return int(math.Max(float64(left), float64(right)))
}