package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	k := scanInt()

	if n < k {
		return
	}

	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = scanInt()
	}

	sort.Ints(h)

	s := 0
	for i := 0; i < n-k; i++ {
		s += h[i]
	}

	fmt.Println(s)
}
