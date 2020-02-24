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

func powUInt2(n int) uint {
	ret := uint(1)
	for i := 0; i < n; i++ {
		ret *= 2
	}
	return ret
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	m := scanInt()

	a := make([]int, m)
	b := make([]uint, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
		nb := scanInt()
		for j := 0; j < nb; j++ {
			b[i] |= powUInt2(scanInt() - 1)
		}
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, powUInt2(n))
	}

	dp[0][0] = 0
	for i := 1; i < int(powUInt2(n)); i++ {
		dp[0][i] = math.MaxInt32
	}

	for i := 0; i < m; i++ {
		for j := uint(0); j < powUInt2(n); j++ {
			l := dp[i][(j&b[i])^j] + a[i]
			r := dp[i][j]
			if l > r {
				dp[i+1][j] = r
			} else {
				dp[i+1][j] = l
			}
		}
	}
	if dp[m][powUInt2(n)-1] == math.MaxInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[m][int(powUInt2(n))-1])
	}
}
