// keep studying(py is a special language,i don't love with it) -- 2020.05.13

package logicfun

import (
	"Common"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 160. Intersection of Two Linked Lists
func GetIntersectionNode(headA, headB *Common.ListNode) *Common.ListNode {
	dict := make(map[*Common.ListNode]struct{})
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := dict[headA]; ok {
				return headA
			} else {
				dict[headA] = struct{}{}
			}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := dict[headB]; ok {
				return headB
			} else {
				dict[headB] = struct{}{}
			}
			headB = headB.Next
		}
	}
	return nil
}

// 162. Find Peak Element
func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}

	res := 1
	for res < len(nums) {
		rcode := 0
		if res == len(nums)-1 {
			rcode = 1
		} else if nums[res] > nums[res+1] {
			rcode = 2
		} else {
			rcode = 3
		}

		if nums[res] > nums[res-1] && (rcode == 1 || rcode == 2) {
			break
		}
		if rcode == 2 {
			res += 2
		} else {
			res++
		}
	}
	return res
}

// 164. Maximum Gap
func MaximumGap(nums []int) int {
	// 3,6,9,1 -> 3
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

// 165. Compare Version Numbers
func CompareVersion(version1 string, version2 string) int {
	// 1.0.0.0001
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0
	for ; i < len(v1) && i < len(v2); i++ {
		vv1, _ := strconv.Atoi(v1[i])
		vv2, _ := strconv.Atoi(v2[i])
		if vv1 > vv2 {
			return 1
		}
		if vv1 < vv2 {
			return -1
		}
	}

	v3 := v1
	res := 1
	if i == len(v1) {
		v3 = v2
		res = -1
	}

	for ; i < len(v3); i++ {
		vv3, _ := strconv.Atoi(v3[i])
		if vv3 != 0 {
			return res
		}
	}

	return 0
}

// 167. Two Sum II - Input array is sorted
func TwoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	left := 1
	right := 2

	for true {
		if numbers[left-1]+numbers[right-1] > target {
			return nil
		}
		if numbers[left-1]+numbers[right-1] == target {
			break
		}
		if right < len(numbers) && (numbers[left-1]+numbers[right] <= target) {
			right++
		} else {
			left++
		}
	}

	return []int{left, right}
}

// 166. Fraction to Recurring Decimal
func FractionToDecimal(numerator int, denominator int) string {
	// assert denominator != 0
	res := ""
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		numerator *= -1
		res += "-"
	}
	red := (numerator % denominator) * 10
	res += strconv.Itoa(numerator / denominator)
	if red != 0 {
		res += "."
	}
	numMap := make(map[int]int)
	index := len(res)

	for red != 0 {
		numMap[red] = index
		quo := red / denominator
		index++
		res += strconv.Itoa(quo)
		red %= denominator
		red *= 10

		if v, ok := numMap[red]; ok {
			res = res[:v] + "(" + res[v:]
			res += ")"
			break
		}
	}

	return res
}

// 168. Excel Sheet Column Title
func ConvertToTitle(n int) string {
	res := ""
	for n > 0 {
		if n == Common.BIG_ENGLISH_CHAR_NUM {
			res = "Z" + res
			return res
		}

		dev := n % Common.BIG_ENGLISH_CHAR_NUM
		if dev > 0 {
			res = string('A'+dev-1) + res
		} else {
			res = "Z" + res
			n -= Common.BIG_ENGLISH_CHAR_NUM
		}
		n /= Common.BIG_ENGLISH_CHAR_NUM
	}
	return res
}

// 169. Majority Element
func MajorityElement(nums []int) int {
	lastNum := nums[0]
	lastAce := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != lastNum {
			lastAce--
			if lastAce == 0 {
				lastNum = nums[i]
				lastAce = 1
			}
		} else {
			lastAce++
		}
	}
	return lastNum
}

// 171. Excel Sheet Column Number
func TitleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= Common.BIG_ENGLISH_CHAR_NUM
		res += int(s[i]-'A') + 1
	}
	return res
}

// 172. Factorial Trailing Zeroes
func TrailingZeroes(n int) int {
	sub := 5
	res := 0
	for n >= sub {
		res += n / sub
		sub *= 5
	}
	return res
}

// 174. Dungeon Game
func CalculateMinimumHP(dungeon [][]int) int {
	type hpHistory struct {
		hp  int
		min int
	}

	n := len(dungeon)
	m := len(dungeon[0])
	dpDungeon := make([][][]hpHistory, n)
	for i := 0; i < n; i++ {
		dpDungeon[i] = make([][]hpHistory, m)
	}
	dpDungeon[0][0] = append(dpDungeon[0][0], hpHistory{hp: dungeon[0][0], min: dungeon[0][0]})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			curDp := dpDungeon[i][j]
			if i > 0 {
				for z := 0; z < len(dpDungeon[i-1][j]); z++ {
					curDp = append(curDp, hpHistory{hp: dpDungeon[i-1][j][z].hp + dungeon[i][j]})
					curDp[len(curDp)-1].min = int(math.Min(float64(curDp[len(curDp)-1].hp), float64(dpDungeon[i-1][j][z].min)))
				}
			}
			if j > 0 {
				for z := 0; z < len(dpDungeon[i][j-1]); z++ {
					isAdd := true
					newDp := hpHistory{
						hp:  dpDungeon[i][j-1][z].hp + dungeon[i][j],
						min: int(math.Min(float64(dpDungeon[i][j-1][z].hp+dungeon[i][j]), float64(dpDungeon[i][j-1][z].min))),
					}

					for s := 0; s < len(curDp); s++ {
						if newDp.hp <= curDp[s].hp && newDp.min <= curDp[s].min {
							isAdd = false
							break
						}
						if newDp.hp >= curDp[s].hp && newDp.min >= curDp[s].min {
							curDp[s].hp = newDp.hp
							curDp[s].min = newDp.min
							isAdd = false
							break
						}
					}
					if isAdd {
						curDp = append(curDp, newDp)
					}
				}
			}
			dpDungeon[i][j] = curDp
		}
	}
	res := Common.MAXINTNUM
	for i := 0; i < len(dpDungeon[n-1][m-1]); i++ {
		if dpDungeon[n-1][m-1][i].min > 0 {
			return 1
		}
		if -dpDungeon[n-1][m-1][i].min+1 < res {
			res = -dpDungeon[n-1][m-1][i].min + 1
		}
	}
	return res
}

// 175. Combine Two Tables # cool: sql question ?
// select FirstName, LastName, City, State from Person left join Address using (PersonId)

// 176. Second Highest Salary
// select IFNULL((select distinct Salary as SecondHighestSalary from Employee order by Salary desc limit 1,1), NULL) as SecondHighestSalary

// 177. Nth Highest Salary
/*
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N = N-1;
  RETURN (
	  select IFNULL((select distinct Salary as SecondHighestSalary from Employee order by Salary desc limit  N,1), NULL)
  );
END
*/

// 178. Rank Scores
/*
SELECT Score, rr as 'Rank' from
(SELECT s1.Score, count(distinct(s2.Score)) as rr
FROM Scores s1,Scores s2
WHERE s1.Score<=s2.Score
GROUP BY s1.Id) as t1
ORDER BY rr
*/

// 179. Largest Number
type dumpNumber struct {
	childs map[int]*dumpNumber
	vals   [10]string
}

func LargestNumber(nums []int) string {
	sortDump := &dumpNumber{childs: make(map[int]*dumpNumber)}
	for i := 0; i < len(nums); i++ {
		curDump := sortDump
		curNum := strconv.Itoa(nums[i])
		for j := 0; j < 10; j++ {
			index := int(curNum[j%len(curNum)] - '0')
			if _, ok := curDump.childs[index]; !ok {
				curDump.childs[index] = &dumpNumber{childs: make(map[int]*dumpNumber)}
			}
			curDump = curDump.childs[index]
		}
		curDump.vals[len(curNum)-1] += curNum
	}

	i := 0
	res := backDump(sortDump, 0)
	for i = 0; i < len(res); i++ {
		if res[i] != '0' {
			break
		}
	}
	if i == len(res) {
		return "0"
	}
	return res[i:]
}

func backDump(dump *dumpNumber, n int) string {
	res := ""
	childs := dump.childs
	if n == 10 {
		// 叶子层
		for i := 0; i <= 9; i++ {
			res = res + dump.vals[i]
		}
	} else {
		for i := 9; i >= 0; i-- {
			if child, ok := childs[i]; ok {
				n++
				res = res + backDump(child, n)
				n--
			}
		}
	}
	return res
}

// 180. Consecutive Numbers
/*
SELECT distinct Num as ConsecutiveNums
from (
  select Num,
    case
      when @prev = Num then @count := @count + 1
      when (@prev := Num) is not null then @count := 1
    end as CNT
  from Logs, (select @prev := null,@count := null) as t
) as temp
where temp.CNT >= 3
*/

// 181. Employees Earning More Than Their Managers
/*
SELECT a.Name as 'Employee'
FROM Employee a,Employee b
WHERE a.ManagerId IS NOT NULL and b.Id = a.ManagerId and a.Salary > b.Salary
*/

// 182. Duplicate Emails
/*
SELECT DISTINCT(a.Email) as 'Email'
FROM Person a, Person b
WHERE a.Email = b.Email and a.Id != B.Id
*/
