package arrays

import "sort"

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

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/

// T: O(N) M: O(N)

func ContainDuplicatesMap(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}

// T: O(N^2) M: O(1)

func ContainDuplicates(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// T: O(N*logN) + O(N) = O(N*logN) M: O(1)

func ContainDuplicatesWithSort(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
