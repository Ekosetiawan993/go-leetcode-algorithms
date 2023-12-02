func twoSum(nums []int, target int) []int {
    leftOver := make(map[int]int)

    for i, n := range(nums) {
        _, ok := leftOver[n]
        if ok {
            return []int{leftOver[n], i}
        }
        leftOver[target-n] = i
    }
        

    return nil
}