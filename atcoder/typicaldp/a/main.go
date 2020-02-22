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
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = scanInt()
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, 100*100+1)
	}

	dp[0][0] = true
	dp[0][p[0]] = true

	for i := 1; i < n; i++ {
		for j := 0; j <= 100*100; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
				dp[i][j+p[i]] = true
			}
		}
	}

	cnt := 0
	for j := 0; j <= 100*100; j++ {
		if dp[n-1][j] {
			cnt++
		}
	}

	fmt.Println(cnt)
}
