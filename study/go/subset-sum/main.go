package main

import "fmt"

func ResolveSubsetSum(as []int, a int) bool {
	// dp[i+1][w]: i番目の要素まで見てwにできるかどうか
	dp := make([][]bool, len(as)+1)
	for i := 0; i <= len(as); i++ {
		dp[i] = make([]bool, a+1)
	}
	dp[0][0] = true

	for i := 0; i < len(as); i++ {
		for w := 0; w <= a; w++ {
			if dp[i][w] {
				dp[i+1][w] = true
				continue
			}
			if w-as[i] >= 0 && dp[i][w-as[i]] {
				dp[i+1][w] = true
				continue
			}
		}
	}
	return dp[len(as)][a]
}

func main() {
	fmt.Println(ResolveSubsetSum([]int{7, 5, 3}, 10))
	fmt.Println(ResolveSubsetSum([]int{9, 7}, 6))
}
