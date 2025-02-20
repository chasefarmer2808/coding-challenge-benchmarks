package leetcode

func RemoveElement(nums []int, val int) int {
	j := 0

	for _, v := range nums {
		if v != val {
			nums[j] = v
			j++
		}
	}

	return j
}
