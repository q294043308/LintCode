package LogicFun

import (
	"Common"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1207. Teemo Attacking        Captain Teemo is on standby
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	sort.Ints(timeSeries)
	lastSerie := timeSeries[0]
	curSerie := timeSeries[0]
	res := 0

	for i := 1; i <= len(timeSeries); i++ {
		if i == len(timeSeries) || curSerie+duration < timeSeries[i] {
			res += curSerie - lastSerie + duration

			if i == len(timeSeries) {
				break
			}

			curSerie = timeSeries[i]
			lastSerie = timeSeries[i]
		} else {
			curSerie = timeSeries[i]
		}
	}
	return res
}

// 1394. Goat Latin
func ToGoatLatin(S string) string {
	if len(S) == 0 {
		return S
	}

	var res string
	appendStr := "maa"
	startIndex := 0
	vowelMap := make(map[byte]struct{})
	vowelMap['a'] = struct{}{}
	vowelMap['e'] = struct{}{}
	vowelMap['i'] = struct{}{}
	vowelMap['o'] = struct{}{}
	vowelMap['u'] = struct{}{}
	vowelMap['A'] = struct{}{}
	vowelMap['E'] = struct{}{}
	vowelMap['I'] = struct{}{}
	vowelMap['O'] = struct{}{}
	vowelMap['U'] = struct{}{}

	for {
		endIndex := strings.Index(S[startIndex:], " ") + startIndex
		var curStr string
		if endIndex < startIndex {
			curStr = S[startIndex:]
			startIndex = -1
		} else {
			curStr = S[startIndex:endIndex]
			startIndex = endIndex + 1
		}

		if len(curStr) == 0 {
			if startIndex == -1 {
				break
			}
			res += " "
			continue
		}

		if _, ok := vowelMap[curStr[0]]; !ok {
			curStr = curStr[1:] + string(curStr[0]) + appendStr
		} else {
			curStr += appendStr
		}
		appendStr += "a"

		res += curStr
		if startIndex != -1 {
			res += " "
		} else {
			break
		}
	}

	return res
}

// 1369. Most Common Word
func MostCommonWord(paragraph string, banned []string) string {
	strCountMap := make(map[string]int)

	for _, str := range banned {
		strCountMap[str] = -1
	}

	startIndex := 0
	var res string
	maxCount := 0

	for {
		endIndexPoint1 := strings.Index(paragraph[startIndex:], ",") + startIndex
		endIndexSpace := strings.Index(paragraph[startIndex:], " ") + startIndex
		endIndexPoint := strings.Index(paragraph[startIndex:], ".") + startIndex
		endIndexaa := strings.Index(paragraph[startIndex:], "!") + startIndex
		if endIndexPoint1 < startIndex {
			endIndexPoint1 = 100000
		}
		if endIndexSpace < startIndex {
			endIndexSpace = 100000
		}
		if endIndexPoint < startIndex {
			endIndexPoint = 100000
		}
		if endIndexaa < startIndex {
			endIndexaa = 100000
		}
		endIndex := int(math.Min(math.Min(math.Min(float64(endIndexPoint1), float64(endIndexSpace)), float64(endIndexPoint)), float64(endIndexaa)))
		var curStr string
		if endIndex < startIndex || endIndex == 100000 {
			curStr = paragraph[startIndex:]
			startIndex = -1
		} else {
			curStr = paragraph[startIndex:endIndex]
			startIndex = endIndex + 1
		}

		if len(curStr) == 0 {
			if startIndex == -1 {
				break
			}
			continue
		}

		curStr = strings.ToLower(curStr)
		if _, ok := strCountMap[curStr]; !ok {
			strCountMap[curStr] = 1
			if maxCount == 0 {
				maxCount = 1
				res = curStr
			}
		} else {
			if strCountMap[curStr] != -1 {
				strCountMap[curStr]++
				if strCountMap[curStr] > maxCount {
					maxCount = strCountMap[curStr]
					res = curStr
				}
			}
		}
		if startIndex == -1 {
			break
		}
	}
	return res
}

// 1201. Next Greater Element II
func NextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	numStack := new(Common.Stack)

	for i := 0; i < 2*len(nums); i++ {
		index := i % len(nums)
		topNum, _ := numStack.Top().(int)
		for !numStack.Empty() && nums[topNum] < nums[index] {
			popNum, _ := numStack.Pop().(int)
			res[popNum] = nums[index]
			topNum, _ = numStack.Top().(int)
		}
		if i < len(nums) {
			res[index] = -1
			numStack.Push(index)
		}
	}
	return res
}

// 1089. Valid Parenthesis String
func CheckValidString(s string) bool {
	left := 0  // 将所有*当做左括号时左括号剩余的个数
	right := 0 // 将所有*当做右括号或空时左括号剩余的个数

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
			right++
		} else {
			if right > 0 {
				right--
			}
			if s[i] == ')' {
				left--
			} else {
				left++
			}
		}
		if left < 0 {
			return false
		}
	}
	return right == 0
}

// 873. Squirrel Simulation
func MinDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	res := 0
	firstSelect := -1
	distanceSelect := float64(0)
	distanceToTreeSelect := float64(0)
	maxSelect := float64(Common.MININTNUM)

	for i := 0; i < len(nuts); i++ {
		toTree := math.Abs(float64(nuts[i][0]-tree[0])) + math.Abs(float64(nuts[i][1]-tree[1]))
		toSquirrel := math.Abs(float64(nuts[i][0]-squirrel[0])) + math.Abs(float64(nuts[i][1]-squirrel[1]))
		res += int(2 * toTree)
		if toTree-toSquirrel > maxSelect {
			firstSelect = i
			maxSelect = toTree - toSquirrel
			distanceSelect = toSquirrel
			distanceToTreeSelect = toTree
		}
	}

	if firstSelect >= 0 {
		res = res - int(distanceToTreeSelect) + int(distanceSelect)
	}
	return res
}

// 1613. Highest frequency IP
func HighestFrequency(ipLines []string) string {
	res := ""
	maxCount := 0
	ipMap := make(map[string]int)

	for _, val := range ipLines {
		ipMap[val]++
		if ipMap[val] > maxCount {
			maxCount = ipMap[val]
			res = val
		}
	}
	return res
}

// 1632. Count email groups
func CountGroups(emails []string) int {
	groupMap := make(map[string]bool)
	res := 0

	for _, val := range emails {
		curVal := ""
		index := 0
		for i, cha := range val {
			if cha == '@' || cha == '+' {
				index = i
				if cha == '+' {
					index = strings.Index(val[index:], "@") + index
				}
				break
			} else if cha == '.' {
				continue
			}
			curVal += string(cha)
		}
		curVal += val[index:]

		_, ok := groupMap[curVal]
		if !ok {
			groupMap[curVal] = false
		} else if !groupMap[curVal] {
			groupMap[curVal] = true
			res++
		}
	}

	return res
}

// 1638. Least Substring
func GetAns(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	curCha := s[0]
	curNum := 1
	res := 1

	for i := 1; i < len(s); i++ {
		if curCha != s[i] || curNum >= k {
			curCha = s[i]
			curNum = 1
			res++
		} else {
			curNum++
		}
	}
	return res
}

// 1623. Minimal Distance In The Array
func MinimalDistance(a []int, b []int) []int {
	trueMap := make(map[int]struct{})
	hashMap := make(map[int]int)
	res := make([]int, len(b))

	for _, val := range a {
		trueMap[val] = struct{}{}
	}

	for index, key := range b {
		if hashVal, ok := hashMap[key]; ok {
			res[index] = hashVal
			continue
		}
		if _, ok := trueMap[key]; ok {
			hashMap[key] = key
			res[index] = key
			continue
		}
		dis := 1
		for true {
			if _, ok := trueMap[key-dis]; ok {
				hashMap[key] = key - dis
				res[index] = key - dis
				break
			}
			if _, ok := trueMap[key+dis]; ok {
				hashMap[key] = key + dis
				res[index] = key + dis
				break
			}
			dis++
		}
	}

	return res
}

//1565. Modern Ludo I
func ModernLudo(length int, connections [][]int) int {
	connectMap := make(map[int]int)
	stepArr := make([]int, length)
	stepArr[0] = 0

	for _, val := range connections {
		connectMap[val[0]] = val[1]
	}

	for i := 1; i < length; i++ {
		if stepArr[i] == 0 {
			if i >= 6 {
				stepArr[i] = stepArr[i-6] + 1
				for j := i - 1; j > i-6; j-- {
					stepArr[i] = int(math.Min(float64(stepArr[i]), float64(stepArr[j]+1)))
				}
			} else {
				stepArr[i] = stepArr[0] + 1
			}
		}

		curindex := i

		for true {
			val, ok := connectMap[curindex+1]
			if !ok {
				break
			}

			stepArr[val-1] = stepArr[curindex]
			if val == length {
				return stepArr[val-1]
			}
			curindex = val - 1
		}
	}

	return stepArr[length-1]
}

//1478. Closest Target Value
func ClosestTargetValue(target int, array []int) int {
	res := array[0] + array[1]
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			sum := array[i] + array[j]
			if (sum > res || res > target) && sum <= target {
				res = sum
			}
		}
	}

	if res > target {
		res = -1
	}
	return res
}

//1510. Buddy Strings
func BuddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	isTrue := 0
	isRepeat := false
	letterArr := make([]bool, Common.SMALL_ENGLISH_CHAR_NUM)
	var exchangeCh byte

	for i := 0; i < len(A); i++ {
		if letterArr[A[i]-'a'] {
			isRepeat = true
		}
		letterArr[A[i]-'a'] = true

		if A[i] != B[i] {
			if isTrue == 0 {
				isTrue = 1
				exchangeCh = A[i]
			} else if isTrue == 1 {
				if exchangeCh == B[i] {
					isTrue = 2
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return isTrue == 2 || isRepeat
}

//1535. To Lower Case
func ToLowerCase(str string) string {
	res := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] <= 'Z' && str[i] >= 'A' {
			res[i] = str[i] - 'A' + 'a'
		} else {
			res[i] = str[i]
		}
	}

	return string(res)
}

func removeOneSub1(matrix [][]int, x int, y int, falgMtx [][]int) bool {
	if x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return false
	}

	if falgMtx[x][y] == 1 {
		return true
	}
	if falgMtx[x][y] == 2 {
		return false
	}

	if matrix[x][y] == 0 {
		return false
	}

	if x == 0 {
		return true
	}

	res := false
	falgMtx[x][y] = 2
	if res = removeOneSub1(matrix, x-1, y, falgMtx) || res; res {
		falgMtx[x][y] = 1
		return res
	}
	if res = removeOneSub1(matrix, x+1, y, falgMtx) || res; res {
		falgMtx[x][y] = 1
		return res
	}
	if res = removeOneSub1(matrix, x, y-1, falgMtx) || res; res {
		falgMtx[x][y] = 1
		return res
	}
	if res = removeOneSub1(matrix, x, y+1, falgMtx) || res; res {
		falgMtx[x][y] = 1
		return res
	}

	falgMtx[x][y] = 2
	return res
}

func RemoveOneSub2(matrix [][]int, x int, y int) {
	if x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return
	}

	if matrix[x][y] == 1 {
		matrix[x][y] = 0
	} else {
		return
	}

	RemoveOneSub2(matrix, x-1, y)
	RemoveOneSub2(matrix, x+1, y)
	RemoveOneSub2(matrix, x, y-1)
	RemoveOneSub2(matrix, x, y+1)
}

// 1621. Cut Connection   想多了----
func RemoveOneOver(matrix [][]int, x int, y int) [][]int {
	if matrix[x][y] == 0 {
		return matrix
	}

	falgMtx := make([][]int, len(matrix))
	for i := 0; i < len(falgMtx); i++ {
		falgMtx[i] = make([]int, len(matrix[0]))
	}

	matrix[x][y] = 0
	falgMtx[x][y] = 2

	if !removeOneSub1(matrix, x-1, y, falgMtx) {
		RemoveOneSub2(matrix, x-1, y)
	}
	if !removeOneSub1(matrix, x+1, y, falgMtx) {
		RemoveOneSub2(matrix, x+1, y)
	}
	if !removeOneSub1(matrix, x, y-1, falgMtx) {
		RemoveOneSub2(matrix, x, y-1)
	}
	if !removeOneSub1(matrix, x, y+1, falgMtx) {
		RemoveOneSub2(matrix, x, y+1)
	}
	return matrix
}

// 1621. Cut Connection
func RemoveOne(matrix [][]int, x int, y int) [][]int {
	for i := x; i < len(matrix); i++ {
		matrix[i][y] = 0
	}

	return matrix
}

// 1779. Shortest Duplicate Subarray
func GetLength(arr []int) int {
	res := -1
	numMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		val, ok := numMap[arr[i]]
		if ok {
			if res == -1 || res > i-val+1 {
				res = i - val + 1
			}
		}
		numMap[arr[i]] = i
	}

	return res
}

func ReachEndpointSub(dataMap [][]int, flag [][]bool, x, y int) bool {
	if x < 0 || y < 0 || x >= len(dataMap) || y >= len(dataMap[0]) || flag[x][y] || dataMap[x][y] == 0 {
		return false
	}
	if dataMap[x][y] == 9 {
		return true
	}

	flag[x][y] = true
	res := false
	if res = ReachEndpointSub(dataMap, flag, x+1, y) || res; res {
		return res
	}
	if res = ReachEndpointSub(dataMap, flag, x, y+1) || res; res {
		return res
	}
	if res = ReachEndpointSub(dataMap, flag, x-1, y) || res; res {
		return res
	}
	if res = ReachEndpointSub(dataMap, flag, x, y-1) || res; res {
		return res
	}

	return res
}

// 1479. Can Reach The Endpoint
func ReachEndpoint(dataMap [][]int) bool {
	flag := make([][]bool, len(dataMap))
	for i := 0; i < len(flag); i++ {
		flag[i] = make([]bool, len(dataMap[0]))
	}

	return ReachEndpointSub(dataMap, flag, 0, 0)
}

// 1615. The result of investment
func GetAns1(funds []int, a int, b int, c int) []int {
	res := []int{a, b, c}

	for _, val := range funds {
		min := res[0]
		minIndex := 0

		for i := 1; i <= 2; i++ {
			if min > res[i] {
				minIndex = i
				min = res[i]
			}
		}

		res[minIndex] += val
	}

	return res
}

//1633. Strings That Satisfies The Condition
func GetAns2(target string, s []string) []string {
	res := make([]string, 0, 1)
	for _, val := range s {
		if len(target) > len(val) {
			continue
		}

		index := 0
		jumpNum := 0
		for _, cha := range val {
			if byte(cha) == target[index] {
				index++
			} else {
				jumpNum++
				if len(target) > len(val)-jumpNum {
					break
				}
			}

			if index == len(target) {
				res = append(res, val)
				break
			}
		}
	}

	return res
}

//1540. Can Convert
func CanConvert(s string, t string) bool {
	jumpNum := 0
	index := 0
	for _, cha := range s {
		if byte(cha) == t[index] {
			index++
		} else {
			jumpNum++
			if len(t) > len(s)-jumpNum {
				break
			}
		}

		if index == len(t) {
			return true
		}
	}

	return false
}

//1645. Least Subsequences
func LeastSubsequences(arrayIn []int) int {
	var res float64 = 0
	resArr := make([]float64, len(arrayIn))

	resArr[0] = 1
	for i := 1; i < len(arrayIn); i++ {
		resArr[i] = 1
		for j := 0; j < i; j++ {
			if arrayIn[i] >= arrayIn[j] {
				resArr[i] = math.Max(resArr[i], resArr[j]+1)
			}
		}
		res = math.Max(resArr[i], res)
	}

	return int(res)
}

// 1480. Dot Product
func DotProduct(A []int, B []int) int {
	if len(A) == 0 || len(A) != len(B) {
		return -1
	}

	res := 0
	for i := 0; i < len(A); i++ {
		res += A[i] * B[i]
	}

	return res
}

//1614. Highest growth stock
func HighestGrowth(Stock [][]string) string {
	var maxNum float64
	res := ""

	for _, val := range Stock {
		v1, _ := strconv.ParseFloat(val[1], 64)
		v2, _ := strconv.ParseFloat(val[2], 64)
		if v2/v1 > maxNum {
			maxNum = v2 / v1
			res = val[0]
		}
	}

	return res
}

func findSubtreeSub(root *Common.TreeNode, maxNum *int, maxNode **Common.TreeNode) int {
	if root == nil {
		return 0
	}

	val := root.Val
	if root.Left != nil {
		val += findSubtreeSub(root.Left, maxNum, maxNode)
	}
	if root.Right != nil {
		val += findSubtreeSub(root.Right, maxNum, maxNode)
	}

	if val > *maxNum {
		*maxNum = val
		*maxNode = root
	}
	return val
}

//628. Maximum Subtree
func FindSubtree(root *Common.TreeNode) *Common.TreeNode {
	if root == nil {
		return nil
	}

	res := new(Common.TreeNode)
	maxNum := Common.MININTNUM

	findSubtreeSub(root, &maxNum, &res)
	return res
}

//956. Data Segmentation
func DataSegmentation(str string) []string {
	startIndex := -1
	endIndex := -1
	res := make([]string, 0, 1)

	for i := 0; i < len(str); i++ {
		if str[i] < 'a' || str[i] > 'z' {
			if startIndex != -1 {
				res = append(res, str[startIndex:endIndex])
			}
			if str[i] != ' ' {
				res = append(res, string(str[i]))
			}

			startIndex = -1
			endIndex = -1
		} else {
			if startIndex == -1 {
				startIndex = i
				endIndex = i + 1
			} else {
				endIndex++
			}
		}
	}

	if startIndex != -1 {
		res = append(res, str[startIndex:endIndex])
	}

	return res
}

//1398. K Decimal Addition
func Addition(k int, a string, b string) string {
	res := ""
	isAdd := false
	i := 0

	for i < len(a) && i < len(b) {
		var curnum int
		curnum = int(a[len(a)-i-1] - '0' + b[len(b)-i-1] - '0')
		if isAdd {
			curnum++
		}

		if curnum >= k {
			curnum -= k
			isAdd = true
		} else {
			isAdd = false
		}

		res = string(curnum+'0') + res
		i++
	}

	remainStr := ""
	if i < len(a) {
		remainStr = a[:len(a)-i]
	}
	if i < len(b) {
		remainStr = b[:len(b)-i]
	}

	if isAdd {
		if remainStr == "" {
			res = "1" + res
		} else {
			res = Addition(k, remainStr, "1") + res
		}
	} else {
		res = remainStr + res
	}

	for res[0] == '0' {
		res = res[1:]
	}

	return res
}

//1617. Array Maximum Difference
func GetAnswer(a []int) int {
	res := 0

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if (a[j] < a[i]) && (a[i]%2 == 0) && (a[j]%2 == 1) && (a[i]-a[j] > res) {
				res = a[i] - a[j]
			}
		}
	}
	return res
}

// 1781. Reverse ASCII Encoded Strings
func ReverseAsciiEncodedString(encodeString string) string {
	res := ""

	for i := 0; i < len(encodeString); i += 2 {
		curNum := (encodeString[i]-'0')*10 + encodeString[i+1] - '0'
		res = string(curNum) + res
	}

	return res
}

// 1410. Matrix Water Injection
func WaterInjection(matrix [][]int, R int, C int) string {
	if R == 0 || R == len(matrix)-1 || C == 0 || C == len(matrix[0])-1 {
		return "YES"
	}

	if R > 0 && matrix[R][C] > matrix[R-1][C] {
		if res := WaterInjection(matrix, R-1, C); res == "YES" {
			return res
		}
	}
	if R < len(matrix)-1 && matrix[R][C] > matrix[R+1][C] {
		if res := WaterInjection(matrix, R+1, C); res == "YES" {
			return res
		}
	}
	if C > 0 && matrix[R][C] > matrix[R][C-1] {
		if res := WaterInjection(matrix, R, C-1); res == "YES" {
			return res
		}
	}
	if C < len(matrix[0])-1 && matrix[R][C] > matrix[R][C+1] {
		if res := WaterInjection(matrix, R, C+1); res == "YES" {
			return res
		}
	}

	return "NO"
}

// 958. Palindrome Data Stream
func GetStream(s string) []int {
	steamMap := make(map[rune]int)
	res := []int{}
	oddCnt := 0

	for _, char := range s {
		_, ok := steamMap[char]
		if !ok {
			steamMap[char] = 0
		}

		steamMap[char]++
		if steamMap[char]%2 == 0 {
			oddCnt--
		} else {
			oddCnt++
		}

		if oddCnt <= 1 {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}

	return res
}

// 1443. Longest AB Substring
func GetAns3(S string) int {
	decNumMap := make(map[int]int)
	decNumMap[0] = 0
	res := 0
	count := 0

	for i := 1; i <= len(S); i++ {
		if S[i-1] == 'A' {
			count++
		} else {
			count--
		}

		val, ok := decNumMap[count]
		if ok {
			if res < i-val {
				res = i - val
			}
		} else {
			decNumMap[count] = i
		}
	}

	return res
}

func findSubtree1Sub(root *Common.TreeNode) (*Common.TreeNode, int, int) {
	var leftNode, rightNode *Common.TreeNode
	var leftVal, rightVal, leftMinVal, rightMinVal int
	leftMinVal = Common.MAXINTNUM
	rightMinVal = Common.MAXINTNUM
	if root.Left != nil {
		leftNode, leftVal, leftMinVal = findSubtree1Sub(root.Left)
	}
	if root.Right != nil {
		rightNode, rightVal, rightMinVal = findSubtree1Sub(root.Right)
	}

	curVal := root.Val + leftVal + rightVal
	if leftMinVal < rightMinVal {
		if curVal < leftMinVal {
			return root, curVal, curVal
		} else {
			return leftNode, curVal, leftMinVal
		}
	} else {
		if curVal < rightMinVal {
			return root, curVal, curVal
		} else {
			return rightNode, curVal, rightMinVal
		}
	}
}

// 596. Minimum Subtree
func FindSubtree1(root *Common.TreeNode) *Common.TreeNode {
	if root == nil {
		return root
	}

	res, _, _ := findSubtree1Sub(root)
	return res
}

func findSubtree2Sub(root *Common.TreeNode) (*Common.TreeNode, int, float64, int) {
	var leftNode, rightNode *Common.TreeNode
	var leftVal, rightVal, leftCnt, rightCnt int
	leftMaxVal := float64(Common.MININTNUM)
	rightMaxVal := float64(Common.MININTNUM)
	if root.Left != nil {
		leftNode, leftVal, leftMaxVal, leftCnt = findSubtree2Sub(root.Left)
	}
	if root.Right != nil {
		rightNode, rightVal, rightMaxVal, rightCnt = findSubtree2Sub(root.Right)
	}

	curCnt := 1 + leftCnt + rightCnt
	curVal := float64(root.Val+leftVal+rightVal) / float64(curCnt)
	if leftMaxVal > rightMaxVal {
		if curVal > leftMaxVal {
			return root, root.Val + leftVal + rightVal, curVal, curCnt
		} else {
			return leftNode, root.Val + leftVal + rightVal, leftMaxVal, curCnt
		}
	} else {
		if curVal > rightMaxVal {
			return root, root.Val + leftVal + rightVal, curVal, curCnt
		} else {
			return rightNode, root.Val + leftVal + rightVal, rightMaxVal, curCnt
		}
	}
}

// 597. Maximum Subtree
func FindSubtree2(root *Common.TreeNode) *Common.TreeNode {
	if root == nil {
		return root
	}

	res, _, _, _ := findSubtree2Sub(root)
	return res
}

// 487. Name Deduplication
func NameDeduplication(names []string) []string {
	resMap := make(map[string]struct{}, 0)
	res := make([]string, 0, 1)

	for _, val := range names {
		resMap[strings.ToLower(val)] = struct{}{}
	}

	for key, _ := range resMap {
		res = append(res, key)
	}

	sort.Strings(res)
	return res
}

// 975. 2 Keys Keyboard
func MinSteps(n int) int {
	minArr := make([]int, n+1)

	for i := 2; i <= n; i++ {
		minArr[i] = i
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				if i/j+minArr[j] < minArr[i] {
					minArr[i] = i/j + minArr[j]
				}
			}
		}
	}

	return minArr[n]
}

// 1298. Minimum Height Trees
func FindMinHeightTrees(n int, edges [][]int) []int {
	res := make([]int, 0, 1)
	edgMap := make(map[int]map[int]struct{})

	if n <= 1 {
		return []int{0}
	}

	for _, edge := range edges {
		if _, ok := edgMap[edge[0]]; !ok {
			edgMap[edge[0]] = make(map[int]struct{})
		}

		edgMap[edge[0]][edge[1]] = struct{}{}

		if _, ok := edgMap[edge[1]]; !ok {
			edgMap[edge[1]] = make(map[int]struct{})
		}
		edgMap[edge[1]][edge[0]] = struct{}{}
	}

	leafNode := make([][]int, 0)
	for len(edgMap) > 2 {
		newLeafNode := make([][]int, 0)
		for _, leaf := range leafNode {
			delete(edgMap[leaf[0]], leaf[1])
			delete(edgMap[leaf[1]], leaf[0])

			if len(edgMap[leaf[0]]) == 0 {
				delete(edgMap, leaf[0])
			} else if len(edgMap[leaf[0]]) == 1 {
				for index, _ := range edgMap[leaf[0]] {
					newLeafNode = append(newLeafNode, []int{leaf[0], index})
				}
			}

			if len(edgMap[leaf[1]]) == 0 {
				delete(edgMap, leaf[1])
			} else if len(edgMap[leaf[1]]) == 1 {
				for index, _ := range edgMap[leaf[1]] {
					newLeafNode = append(newLeafNode, []int{leaf[1], index})
				}
			}
		}

		if len(newLeafNode) == 0 {
			for index, _ := range edgMap {
				if len(edgMap[index]) == 1 {
					for in, _ := range edgMap[index] {
						newLeafNode = append(newLeafNode, []int{index, in})
					}
				}

			}
		}

		leafNode = newLeafNode
	}

	for index, _ := range edgMap {
		res = append(res, index)
	}
	if len(res) == 0 {
		for _, Node := range leafNode {
			res = append(res, Node[0])
		}
	}

	return res
}
