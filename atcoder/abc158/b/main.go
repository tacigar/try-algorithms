package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt64() int64 {
	scanner.Scan()
	a, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt64()
	a := scanInt64()
	b := scanInt64()
	if n < a+b {
		if n < a {
			fmt.Println(n)
		} else {
			fmt.Println(a)
		}
	} else {
		x := n / (a + b)
		y := n - x*(a+b)
		if y > a {
			fmt.Println(x*a + a)
		} else {
			fmt.Println(x*a + y)
		}
	}
}
