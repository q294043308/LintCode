package Common

// 究竟是什么白痴会把数组的判空做成这样？
func LintCodeLen(nums []int) int {
	if len(nums) == 1 && nums[0] == 0 {
		return 0
	}

	return len(nums)
}
