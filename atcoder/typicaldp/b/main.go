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

	na := scanInt()
	nb := scanInt()

	a := make([]int, na)
	for i := 0; i < na; i++ {
		a[i] = scanInt()
	}

	b := make([]int, nb)
	for i := 0; i < nb; i++ {
		b[i] = scanInt()
	}

	dp := make([][]int, na+1)
	for i := range dp {
		dp[i] = make([]int, nb+1)
	}

	dp[na][nb] = 0

	for i := na; i >= 0; i-- {
		for j := nb; j >= 0; j-- {
			if i == na && j == nb {
				continue
			}
			if (i+j)%2 == 0 {
				if i == na {
					dp[i][j] = dp[i][j+1] + b[j]
				} else if j == nb {
					dp[i][j] = dp[i+1][j] + a[i]
				} else {
					r := dp[i+1][j] + a[i]
					l := dp[i][j+1] + b[j]
					if r > l {
						dp[i][j] = r
					} else {
						dp[i][j] = l
					}
				}
			} else {
				if i == na {
					dp[i][j] = dp[i][j+1]
				} else if j == nb {
					dp[i][j] = dp[i+1][j]
				} else {
					r := dp[i+1][j]
					l := dp[i][j+1]
					if r > l {
						dp[i][j] = l
					} else {
						dp[i][j] = r
					}
				}
			}
		}
	}
	fmt.Println(dp[0][0])
}
