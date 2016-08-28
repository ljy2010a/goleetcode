package leetcode

import (
	"fmt"
	"sort"
)

func threesum(nums []int) [][]int {
	var ansList [][]int
	length := len(nums)
	sort.Ints(nums)
	fmt.Printf("%v \n", nums)

	last_head := 1 << 32

	for i := 0; i < length - 2; i++ {

		low, high := i + 1, length - 1
		sum := 0 - nums[i]
		last_ans := []int{}
		if last_head == nums[i] {
			continue
		}

		last_head = nums[i]

		for {

			if low >= high {
				break
			}

			switch {
			case nums[low] + nums[high] == sum:
				if 3 == len(last_ans) && (last_ans[1] == nums[low] || last_ans[1] == nums[high] ) {
					low++
					high = high - 1
					continue
				}
				var ans []int = []int{nums[i], nums[low], nums[high]}
				fmt.Println(ans)
				ansList = append(ansList, ans)
				last_ans = ans
				low++
				high = high - 1

			case nums[low] + nums[high] < sum:
				low++
			case nums[low] + nums[high] > sum:
				high = high - 1
			}

		}
	}
	return ansList
}
