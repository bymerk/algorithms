package arrays

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/

// T: O(N) M: O(N)

func RotateWithAllocate(nums []int, k int) []int {
	n := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n[(i+k)%len(nums)] = nums[i]
	}

	for i := 0; i < k; i++ {
		n[i] = nums[len(nums)-i-1]
	}

	return n
}

func Rotate(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			nums[j], prev = prev, nums[j]
		}
	}
	return nums
}
