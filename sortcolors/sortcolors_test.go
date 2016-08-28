package leetcode

import (
	"testing"
)

func TestSort_colors(t *testing.T) {
	data := []int{1, 1, 1, 2, 0, 2, 2, 0, 1}
	sort_colors(data)
	t.Log(data)
}
