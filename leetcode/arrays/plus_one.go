package arrays

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/

func PlusOne(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var overflow bool
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 9 {
			nums[i]++
			break
		}

		nums[i] = 0
		if i == 0 {
			overflow = true
		}
	}

	if overflow {
		n := make([]int, 0, len(nums)+1)
		n = append(n, 1)
		return append(n, nums...)
	}

	return nums
}
