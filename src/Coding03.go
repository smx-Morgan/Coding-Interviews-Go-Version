package src

func findRepeatNumber(nums []int) int {
	Hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if Hashmap[nums[i]] == 1 {
			return nums[i]
		}
		if Hashmap[nums[i]] == 0 {
			Hashmap[nums[i]] = 1
		}
	}
	return -1
}
