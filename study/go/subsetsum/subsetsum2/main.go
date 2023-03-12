package main

import "fmt"

func ResolveSubsetSum(as []int, sum int) bool {
	// dp[i+1][s]: iまで見た時にsの部分和を作れるかどうか
	n := len(as)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for s := 0; s <= sum; s++ {
			if dp[i][s] {
				dp[i+1][s] = dp[i][s]
			} else if s-as[i] >= 0 {
				dp[i+1][s] = dp[i][s-as[i]]
			}
		}
	}
	return dp[n][sum]
}

func main() {
	fmt.Println(ResolveSubsetSum([]int{7, 5, 3}, 10))
	fmt.Println(ResolveSubsetSum([]int{9, 7}, 6))
}
