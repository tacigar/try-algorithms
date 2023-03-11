package main

import "fmt"

func ResolveSubsetSum(vs []int, sum int) bool {
	n := len(vs)

	// dp[i+1][s]: iまで見た時にsにできるかどうか
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for s := 0; s <= sum; s++ {
			if s-vs[i] >= 0 {
				dp[i+1][s] = dp[i][s] || dp[i][s-vs[i]]
			}
		}
	}
	return dp[n][sum]
}

func main() {
	fmt.Println(ResolveSubsetSum([]int{7, 5, 3}, 10))
	fmt.Println(ResolveSubsetSum([]int{9, 7}, 6))
}
