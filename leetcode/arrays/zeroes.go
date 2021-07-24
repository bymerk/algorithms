package arrays

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/

func MoveZeroes(nums []int) []int {
	var positiveCount int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[positiveCount] = nums[i]
			positiveCount++
		}
	}

	for i := positiveCount; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func MoveZeroesWithAllocation(nums []int) []int {
	n := make([]int, len(nums))
	var count int
	for _, v := range nums {
		if v != 0 {
			n[count] = v
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		n[i] = 0
	}

	return n
}
