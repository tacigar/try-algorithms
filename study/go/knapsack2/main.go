package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func ResolveKnapsack(weight, value []int, maxW int) (int, error) {
	if len(weight) != len(value) {
		return 0, errors.New("invalid input")
	}
	n := len(weight)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxW+1)
	}

	for i := 0; i < n; i++ {
		for w := 0; w <= maxW; w++ {
			v := dp[i][w]
			if w-weight[i] >= 0 {
				v = int(math.Max(float64(v), float64(dp[i][w-weight[i]]+value[i])))
			}
			dp[i+1][w] = v
		}
	}

	return dp[n][maxW], nil
}

func main() {
	r, err := ResolveKnapsack([]int{2, 1, 3, 2, 1, 5}, []int{3, 2, 6, 1, 3, 85}, 9)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(r)
}
