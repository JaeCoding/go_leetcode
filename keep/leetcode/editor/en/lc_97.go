package main

//Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
//
//
// Example 1:
//
//
//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//Output: true
//
//
// Example 2:
//
//
//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//Output: false
//
// Related Topics String Dynamic Programming

//leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {

	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true
	for i := 0; i < len(dp); i++ {

		for j := 0; j < len(dp[0]); j++ {
			// 看左 和上 是否为T 为T可以考虑计算，
			if i > 0 && dp[i-1][j] {
				if s1[i-1:i] == s3[i+j-1:i+j] {
					dp[i][j] = true
					continue
				}
			}
			// 看左
			if j > 0 && dp[i][j-1] {
				if s2[j-1:j] == s3[i+j-1:i+j] {
					dp[i][j] = true
					continue
				}
			}
			if i == 0 && j == 0 {
				continue
			}

			dp[i][j] = false
		}

	}

	return dp[l1][l2]
}

func main() {
	//isInterleave("aabcc", "dbbca", "aadbbcbcac")
	println(isInterleave("a", "", "a"))
}

//leetcode submit region end(Prohibit modification and deletion)
