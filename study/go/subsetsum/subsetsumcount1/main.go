package main

import "fmt"

func ResolveSubsetSumCount(as []int, sum int) int {
	n := len(as)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
	}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for w := 0; w <= sum; w++ {
			dp[i+1][w] = dp[i][w]
			if w-as[i] >= 0 {
				dp[i+1][w] += dp[i][w-as[i]]
			}
		}
	}

	return dp[n][sum]
}

func main() {
	fmt.Println(ResolveSubsetSumCount([]int{7, 5, 3, 1, 8}, 12))
	fmt.Println(ResolveSubsetSumCount([]int{4, 1, 1, 1}, 5))
}
