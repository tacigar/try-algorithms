package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	as := make([][]int, 2)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		as[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			as[i][j] = scanInt()
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = dp[i][j] + as[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + as[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + as[i][j]
			} else {
				dp[i][j] = max(
					dp[i][j-1]+as[i][j],
					dp[i-1][j]+as[i][j],
				)
			}
		}
	}
	fmt.Println(dp[1][n-1])
}
