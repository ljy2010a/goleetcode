package leetcode

import "fmt"

func sort_colors(nums []int) {

	left, mid := 0, 0
	right := len(nums) - 1
	fmt.Printf("%v \n", nums)
	for {
		if mid > right {
			break
		}
		switch nums[mid] {
		case 0:
			nums[mid], nums[left] = nums[left], nums[mid]
			left += 1
			mid += 1
		case 1:
			mid += 1
		case 2:
			nums[mid], nums[right] = nums[right], nums[mid]
			right -= 1
		default:
			break
		}
	}
	fmt.Printf("%v \n", nums)

}
