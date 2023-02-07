package benchmark

func TwoEqualToTargetLine(nums []int, target int) bool {
	mp := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[target-nums[i]]; ok {
			return true
		}
		mp[nums[i]] = struct{}{}
	}
	return false
}

func TwoEqualToTargetSquared(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}
