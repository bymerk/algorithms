package arrays

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/

// T: O(N) M: O(1)

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prev := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[count], prev = prev, nums[i]
			count++
		}
	}

	return len(nums[0:count])
}
