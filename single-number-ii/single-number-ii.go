package leetcode

import "fmt"

func singleNumber(nums []int) int {
	var ones, twos, threes = 0, 0, 0
	for i := 0; i < len(nums); i++ {
		twos |= (ones & nums[i])
		ones ^= nums[i]
		threes = ^(ones & twos)
		ones &= threes
		twos &= threes
		fmt.Printf("ones : %v \n", ones)
		fmt.Printf("twos : %v \n", twos)
		fmt.Printf("threes : %v \n", threes)
		fmt.Printf("======\n")
	}
	return ones
}
