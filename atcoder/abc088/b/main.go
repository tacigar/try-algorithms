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
	value, _ := strconv.Atoi(scanner.Text())
	return value
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scanInt()
	}
	sort.Ints(as)
	alice := 0
	bob := 0
	for i := 0; i < n; i++ {
		idx := n - i - 1
		if i%2 == 0 {
			alice += as[idx]
		} else {
			bob += as[idx]
		}
	}
	fmt.Println(alice - bob)
}
