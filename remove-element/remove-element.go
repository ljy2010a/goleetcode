package leetcode

func removeElement(nums []int, val int) int {
	l := len(nums)
	s := 0
	for i := 0; i < l; i++ {
		if nums[i] == val {
			continue
		}
		nums[s] = nums[i]
		s++
	}
	return s
}
