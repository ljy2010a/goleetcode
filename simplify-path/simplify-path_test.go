package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(simplifyPath("/home/"))
	t.Log(simplifyPath("/a/./b/../../c/"))
	t.Log(simplifyPath("/../c/"))
	t.Log(simplifyPath("/../"))
}
