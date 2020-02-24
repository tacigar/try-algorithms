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

	scanInt()
	m := scanInt()

	l := scanInt() - 1
	r := scanInt() - 1
	for i := 1; i < m; i++ {
		ll := scanInt() - 1
		rr := scanInt() - 1
		if ll > l {
			l = ll
		}
		if rr < r {
			r = rr
		}
	}
	if l > r {
		fmt.Println(0)
	} else {
		fmt.Println(r - l + 1)
	}
}
