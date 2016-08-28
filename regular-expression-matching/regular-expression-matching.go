package leetcode

import "fmt"

func isMatch(s string, p string) bool {

	sLen := len(s)
	pLen := len(p)

	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	for i := 0; i < sLen+1; i++ {
		for j := 1; j < pLen+1; j++ {
			if p[j-1] != '.' && p[j-1] != '*' {
				if i > 0 && s[i-1] == p[j-1] && dp[i-1][j-1] {
					dp[i][j] = true
				}
			} else if p[j-1] == '.' {
				if i > 0 && dp[i-1][j-1] {
					dp[i][j] = true
				}
			} else if j > 1 {
				if dp[i][j-1] || dp[i][j-2] {
					dp[i][j] = true
				} else if i > 0 && (p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j] {
					dp[i][j] = true
				}
			}
		}
	}

	for i := 0; i < sLen+1; i++ {
		fmt.Println(dp[i])
	}
	return dp[sLen][pLen]
}
