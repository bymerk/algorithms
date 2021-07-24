package arrays

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/

func Intersection(nums1, nums2 []int) []int {
	var intersect = make(map[int]int, len(nums1))
	for _, v := range nums1 {
		intersect[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := intersect[v]; ok {
			intersect[v] = 2
		}
	}

	var elements []int
	for k, v := range intersect {
		if v == 2 {
			elements = append(elements, k)
		}
	}

	return elements
}
