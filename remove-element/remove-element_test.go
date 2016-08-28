package leetcode

import (
	"log"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{22, 22, 33}
	val := 33
	ret := removeElement(nums, val)
	log.Println(ret)
}
