package leetcode

import "testing"

func Test(t *testing.T) {
	s := "zabbaxaaa"
	p := "z*a.bbax*a"

	// s := "aab"
	// p := "a*b"
	// s := "aab"
	// p := "c*a*b"
	t.Log(isMatch(s, p))
}
