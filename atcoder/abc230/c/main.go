package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	a := scanInt()
	b := scanInt()
	p := scanInt()
	q := scanInt()
	r := scanInt()
	s := scanInt()
	min1 := max(1-a, 1-b)
	min2 := max(1-a, b-n)
	max1 := min(n-a, n-b)
	max2 := min(n-a, b-1)

	var sb strings.Builder
	for i := p; i <= q; i++ {
		for j := r; j <= s; j++ {
			if i-a == j-b && min1 <= i-a && i-a <= max1 {
				sb.WriteRune('#')
			} else if i-a == b-j && min2 <= i-a && i-a <= max2 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Print(sb.String())
}
