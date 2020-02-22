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

	l := 0
	r := 40

	for r > l+1 {
		c := (l + r) / 2
		if float64(h) >= math.Pow(2, float64(c)) {
			l = c
		} else {
			r = c
		}
	}
	x := l

	ret := make([]int, x+1)
	ret[0] = 1
	for i := 1; i <= x; i++ {
		ret[i] = ret[i-1]*2 + 1
	}
	fmt.Println(ret[x])
}
