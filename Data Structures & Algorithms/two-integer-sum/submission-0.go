func twoSum(nums []int, t int) []int {
    resMap := make(map[int]int)

	for i, v := range nums {
		k := t - v

		if j, ok := resMap[k]; ok {
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		}

		resMap[v] = i
	}

	return []int{}
}
