package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		t := target - num
		if _, ok := m[t]; ok {
			return []int{m[t], i}
		}
		m[num] = i
	}
	return nil
}
