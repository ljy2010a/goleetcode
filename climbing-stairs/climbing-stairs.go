package leetcode

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	a := 2
	b := 3
	c := 5
	for i := 5; i <= n; i++ {
		a = c
		c = b + c
		b = a
	}
	return c
}
