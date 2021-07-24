package sort

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	middle := len(nums) / 2
	return merge(MergeSort(nums[0:middle]), MergeSort(nums[middle:]))
}

func merge(a, b []int) []int {
	combined := make([]int, 0, len(a)+len(b))
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			combined = append(combined, a[i])
			i++
		} else {
			combined = append(combined, b[j])
			j++
		}
	}

	for i < len(a) {
		combined = append(combined, a[i])
		i++
	}

	for j < len(b) {
		combined = append(combined, b[j])
		j++
	}

	return combined
}

func MergeSortWithoutAppend(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	middle := len(nums) / 2
	return mergeWithoutAppend(
		MergeSortWithoutAppend(nums[0:middle]),
		MergeSortWithoutAppend(nums[middle:]),
	)
}

func mergeWithoutAppend(a, b []int) []int {
	combined := make([]int, len(a)+len(b))
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			combined[i+j] = a[i]
			i++
		} else {
			combined[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		combined[i+j] = a[i]
		i++
	}

	for j < len(b) {
		combined[i+j] = b[j]
		j++
	}

	return combined
}
