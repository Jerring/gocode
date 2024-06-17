package leetcode

func findMiddleIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left := 0
	for i, num := range nums {
		if left == sum - left - num {
			return i
		}
		left += num
	}
	return -1
}