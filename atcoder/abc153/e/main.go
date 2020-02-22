package main

import (
	"bufio"
	"fmt"
	"math"
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

	h := scanInt()
	n := scanInt()

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
	}

	dp := make([]int, h+1)
	dp[0] = 0

	for i := 1; i <= h; i++ {
		min := math.MaxInt32
		for j := 0; j < n; j++ {
			if a[j] <= i {
				v := b[j] + dp[i-a[j]]
				if v < min {
					min = v
				}
			} else {
				if b[j] < min {
					min = b[j]
				}
			}
		}
		dp[i] = min
	}
	fmt.Println(dp[h])
}
