package leetcode

import "math"

func minCut(s string) int {
	if s == "" {
		return 0
	}

	s_len := len(s)

	D := make([]int, s_len+1)
	D[0] = -1

	P := make([][]bool, s_len)
	for i := range P {
		P[i] = make([]bool, s_len)
	}
	for i := 1; i <= s_len; i++ {
		D[i] = i - 1
		for j := 0; j <= i-1; j++ {
			P[j][i-1] = false
			if s[j] == s[i-1] && (i-1-j <= 2 || P[j+1][i-2]) {
				P[j][i-1] = true
				D[i] = int(math.Min(float64(D[i]), float64(D[j]+1)))
			}
		}
	}

	return D[s_len]
}
