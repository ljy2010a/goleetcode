package leetcode

import "log"

func subsets(nums []int) [][]int {
	sets := [][]int{}
	cur := []int{}
	nLen := len(nums)
	// sort.Sort(sort.IntSlice(nums))
	log.Println(nums)
	sets = append(sets, []int{})
	for i := 0; i < nLen; i++ {
		sLen := len(sets)
		for j := 0; j < sLen; j++ {
			// cur = sets[j]
			cur = make([]int, len(sets[j]))
			copy(cur, sets[j])
			log.Println(cur, "+", nums[i])
			cur = append(cur, nums[i])
			log.Println("->", cur)
			sets = append(sets, cur)
		}
		log.Println(sets)
	}
	return sets
}
