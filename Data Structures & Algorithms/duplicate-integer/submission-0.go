func hasDuplicate(nums []int) bool {
    res := make(map[int]int)

    for _, v := range(nums) {
        res[v]++
    }

    for _, v := range(res) {
        if v > 1 {
            return true
        }
    }

    return false
}
