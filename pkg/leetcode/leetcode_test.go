package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := RemoveElement(nums, val)

	if k != 2 {
		t.Fatalf("expected to remove 2 elements but removed %d", k)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	k := RemoveDuplicates(nums)

	if k != 2 {
		t.Fatalf("expected 2 unique elements but got %d", k)
	}
}
