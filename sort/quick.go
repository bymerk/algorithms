package sort

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	q := partition(nums)
	QuickSort(nums[:q])
	QuickSort(nums[q+1:])
	return nums
}

func partition(nums []int) int {
	r := len(nums) - 1
	last := nums[r]
	i := -1
	for j := 0; j < r; j++ {
		if nums[j] <= last {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
