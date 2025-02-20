package leetcode

func RemoveDuplicates(nums []int) int {
	seen := make(map[int]bool)

	j := 0
	for _, n := range nums {
		_, ok := seen[n]
		if !ok {
			seen[n] = true
			nums[j] = n
			j++
		}
	}

	return len(seen)
}
