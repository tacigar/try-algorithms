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
	t := scanner.Text()
	n, _ := strconv.Atoi(t)
	return n
}

type C struct {
	a int
	b int
	c int
}

var dp [2001][20001]bool

func fn(n, y int, c C) (bool, *C) {
	if y < 0 || dp[n][y/1000] {
		return false, nil
	}
	if n == 0 {
		if y == 0 {
			return true, &c
		}
		dp[n][y/1000] = true
		return false, nil
	}

	r1, c1 := fn(n-1, y-10000, C{c.a + 1, c.b, c.c})
	if r1 {
		return true, c1
	}
	r2, c2 := fn(n-1, y-5000, C{c.a, c.b + 1, c.c})
	if r2 {
		return true, c2
	}
	r3, c3 := fn(n-1, y-1000, C{c.a, c.b, c.c + 1})
	if r3 {
		return true, c3
	}
	dp[n][y/1000] = true
	return false, nil
}

func main() {
	scanner.Split(bufio.ScanWords)
	_, c := fn(scanInt(), scanInt(), C{0, 0, 0})
	if c != nil {
		fmt.Printf("%d %d %d\n", c.a, c.b, c.c)
	} else {
		fmt.Println("-1 -1 -1")
	}
}
