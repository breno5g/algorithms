package leetcode

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i, v := range nums {
		if j, ok := table[target-v]; ok {
			return []int{j, i}
		}
		table[v] = i
	}
	return nil
}
